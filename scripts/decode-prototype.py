#!/usr/bin/env python3
"""Decode gzip+base64 prototype sources into prototype/."""
from pathlib import Path
import base64, gzip, sys

ROOT = Path(__file__).resolve().parents[1]
ENC = ROOT / "prototype" / "encoded"
OUT = ROOT / "prototype"

names = [
    "index.html",
    "styles.css",
    "app.js",
    "clinic-mobile-single.html",
]

def main():
    ENC.mkdir(parents=True, exist_ok=True)
    for name in names:
        p = ENC / f"{name}.gz.b64"
        if not p.exists():
            print("missing", p, file=sys.stderr)
            continue
        raw = gzip.decompress(base64.b64decode(p.read_text().encode("ascii")))
        out = OUT / name
        out.write_bytes(raw)
        print("wrote", out, len(raw), "bytes")
    single = OUT / "clinic-mobile-single.html"
    if single.exists():
        (OUT / "clinic-mobile.html").write_bytes(single.read_bytes())
        print("wrote", OUT / "clinic-mobile.html")

if __name__ == "__main__":
    main()
