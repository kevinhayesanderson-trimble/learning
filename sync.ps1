param(
    [string]$Message
)

$repoRoot = $PSScriptRoot
Set-Location $repoRoot

Write-Host "==> Checking repository status..." -ForegroundColor Cyan

# Fetch latest remote references
git fetch origin

# Check if there are local working directory changes
$hasChanges = (git status --porcelain)

if ($hasChanges) {
    if (-not $Message) {
        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $computer = $env:COMPUTERNAME
        $Message = "Sync learnings [$computer] at $timestamp"
    }
    Write-Host "==> Staging and committing local changes: '$Message'..." -ForegroundColor Yellow
    git add -A
    git commit -m "$Message"
} else {
    Write-Host "==> No uncommitted local changes." -ForegroundColor Green
}

# Pull latest changes from remote with rebase
Write-Host "==> Pulling latest changes from origin/main..." -ForegroundColor Cyan
git pull --rebase --autostash origin main

if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Conflict during pull/rebase. Please resolve conflicts and run 'git rebase --continue'." -ForegroundColor Red
    exit 1
}

# Push local commits to remote
Write-Host "==> Pushing to origin/main..." -ForegroundColor Cyan
git push origin main

if ($LASTEXITCODE -eq 0) {
    Write-Host "[SUCCESS] Synchronized successfully with GitHub!" -ForegroundColor Green
} else {
    Write-Host "[ERROR] Push failed. Check your permissions or network connection." -ForegroundColor Red
    exit $LASTEXITCODE
}
