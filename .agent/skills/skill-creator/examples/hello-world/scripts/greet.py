#!/usr/bin/env python3
import argparse

def main():
    parser = argparse.ArgumentParser(description="Greet the user.")
    parser.add_argument("--name", default="World", help="The name to greet")
    args = parser.parse_args()
    
    print(f"👋 Hello, {args.name}! Welcome to Antigravity.")

if __name__ == "__main__":
    main()
