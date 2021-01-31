'''
'''
import pprint
from flask import Blueprint, request
from web.common import jsonify_dataframe
import akshare as ak
fund = Blueprint('basic', __name__, url_prefix='/fund')


@fund.route('/stock_fund_hold', methods=['GET'])
def stock_fund_hold():
    stockCode = request.json.get('stockCode')
    stock_fund_stock_holder_df = ak.stock_fund_stock_holder(stock=stockCode)
    stock_fund_stock_holder_df.columns = [
        "fundAbbr", "fundCode", "volume", "percetage", "marketValue", "date"
    ]
    return jsonify_dataframe(stock_fund_stock_holder_df)


