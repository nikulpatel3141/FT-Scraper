# FT Market Data Scraper

The `main.py` script queries an endpoint provided by [FT](https://www.ft.com/) for the daily returns of the instruments (hardcoded) in that script.

The `main.go` script does the same thing (but is written in Golang instead of Python).

## Automated Scraping

The `scrape.yml` workflow file runs the script every 5 minutes and pushes the changes to the `master` branch.

Adapted from this tutorial: https://www.swyx.io/github-scraping

## Copyright

As far as I can tell from the FT's copyright page this isn't not breaking any rules. If it is let me know and I'll take it down.

Ref: https://help.ft.com/legal-privacy/copyright-policy/
