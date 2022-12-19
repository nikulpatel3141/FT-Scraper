"""Query FT for daily returns for the instruments specified in symbols.json
and print the output to the console
"""

import json

import requests

FT_SYMBOLS = {
    "FTSE100": "FTSE:FSI",
    "DJIA": "DJI:DJI",
    "GBPUSD": "GBPUSD",
    "Coffee": "KC.1:IUS",
    "Gold": "GC.1:CMX",
}  # FIXME: should be in a config file but more convenient here
FT_ENDPOINT = (
    "https://markets-data-api-proxy.ft.com/research/webservices/securities/v1/quotes"
)


def query_ft_mkt_data(symbols: list) -> dict:
    """Query FT for market data quotes

    Returns: a dict of quote data with the queried symbols as keys
    """
    endpoint = FT_ENDPOINT + "?symbols=" + ",".join(symbols)
    data = requests.get(endpoint).json()
    data = data["data"]["items"]
    return {
        sym: quote["quote"]
        for quote in data
        for sym in symbols
        if quote["symbolInput"] == sym
    }


def main():
    ft_data = query_ft_mkt_data(list(FT_SYMBOLS.values()))
    daily_ret = {
        disp_sym: ft_data[ft_sym]["change1DayPercent"]
        for disp_sym, ft_sym in FT_SYMBOLS.items()
    }

    daily_ret_json = json.dumps(daily_ret)
    print(daily_ret_json)


if __name__ == "__main__":
    main()
