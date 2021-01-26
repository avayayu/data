'''

'''
from flask import Blueprint
from web.common import jsonify_dataframe, jsonify_failure
import akshare as ak
basic = Blueprint('basic', __name__, url_prefix='/basic')


@basic.route('/stock_info_a_code_name', methods=['GET'])
def stock_info_a_code_name():
    try:
        stock_info_a_code_name_df = ak.stock_info_a_code_name()
        return jsonify_dataframe(stock_info_a_code_name_df)
    except Exception as e:
        jsonify_failure(e)
