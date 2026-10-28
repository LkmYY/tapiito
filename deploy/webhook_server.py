import os
import hmac
import hashlib
import subprocess
from fastapi import FastAPI, Request, HTTPException


app = FastAPI()


def _valid_signature(secret: str, token: str, body: bytes) -> bool:
    if not secret:
        return False
    if token == secret:
        return True
    try:
        sig = hmac.new(secret.encode("utf-8"), body, hashlib.sha256).hexdigest()
        return hmac.compare_digest(sig, token)
    except Exception:
        return False


@app.post("/webhook/gitee")
async def gitee_webhook(request: Request):
    body = await request.body()
    token = request.headers.get("X-Gitee-Token", "")
    event = request.headers.get("X-Gitee-Event", "")
    secret = os.getenv("WEBHOOK_SECRET", "").strip()
    if not _valid_signature(secret, token, body):
        raise HTTPException(status_code=403, detail="invalid signature")

    payload = await request.json()
    ref = str(payload.get("ref", ""))
    branch = os.getenv("BRANCH", "main")
    expected_ref = f"refs/heads/{branch}"
    if ref != expected_ref:
        return {"ok": True, "skipped": True, "reason": "branch mismatch", "ref": ref}

    repo_dir = os.getenv("REPO_DIR", os.getcwd())
    deploy_script = os.getenv("DEPLOY_SCRIPT", os.path.join(repo_dir, "deploy", "deploy.sh"))
    if not os.path.exists(deploy_script):
        raise HTTPException(status_code=500, detail="deploy script missing")

    try:
        proc = subprocess.run(["bash", deploy_script], cwd=repo_dir, capture_output=True, text=True)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"run failed: {e}")
    if proc.returncode != 0:
        raise HTTPException(status_code=500, detail={"stdout": proc.stdout[-2000:], "stderr": proc.stderr[-2000:]})
    return {"ok": True, "event": event, "stdout_tail": proc.stdout[-2000:]}

