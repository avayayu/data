'''
'''
import pprint
from flask import Blueprint, request
from web.common import jsonify_dataframe
import akshare as ak
share = Blueprint('share', __name__, url_prefix='/share')


@share.route('/stock_main_stock_holder', methods=['GET'])
def stock_main_stock_holder():

    stockCode = request.GET.get('stockCode')
    stock_main_stock_holder_df = ak.stock_main_stock_holder(stock=stockCode)
    stock_main_stock_holder_df["stockCode"] = stockCode
    stock_main_stock_holder_df.columns = [[
        "stockCode", "股东名称", "持股数量(股)", "持股比例(%)", "股本性质", "截至日期", "公告日期",
        "股东说明", "股东总数", "平均持股数"
    ]]
