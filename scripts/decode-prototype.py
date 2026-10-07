#!/usr/bin/env python3
"""Decode chunked gzip+base64 prototype sources into prototype/."""
from pathlib import Path
import base64, gzip, sys

ROOT = Path(__file__).resolve().parents[1]
ENC = ROOT / "prototype" / "encoded"
OUT = ROOT / "prototype"
NAMES = ["index.html", "styles.css", "app.js", "clinic-mobile-single.html"]

def load_b64(name: str) -> str:
    single = ENC / f"{name}.gz.b64"
    if single.exists() and single.stat().st_size > 200:
        return single.read_text().strip()
    parts = sorted(ENC.glob(f"{name}.gz.b64.part*"))
    if not parts:
        raise FileNotFoundError(name)
    return "".join(p.read_text().strip() for p in parts)

def main():
    OUT.mkdir(parents=True, exist_ok=True)
    for name in NAMES:
        try:
            b64 = load_b64(name)
        except FileNotFoundError:
            print("missing", name, file=sys.stderr)
            continue
        raw = gzip.decompress(base64.b64decode(b64.encode("ascii")))
        (OUT / name).write_bytes(raw)
        print("wrote", name, len(raw))
    single = OUT / "clinic-mobile-single.html"
    if single.exists():
        (OUT / "clinic-mobile.html").write_bytes(single.read_bytes())

if __name__ == "__main__":
    main()
