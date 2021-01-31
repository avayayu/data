'''
'''
import pprint
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


@basic.route('/stock_info_all', methods=['GET'])
def stock_info_all():
    stock_info_sh_df = ak.stock_info_sh_name_code(indicator="主板A股")
    df = stock_info_sh_df[['company_abbr', 'stockCode', 'list_date']]
    stock_info_sz_df = ak.stock_info_sz_name_code(indicator="A股列表")[[
        'A股简称', 'A股代码', 'A股上市日期'
    ]]
    stock_info_sz_df.columns = ['companyAbbr', 'stockCode', 'listDate']

    df = df.append(stock_info_sz_df)
    return jsonify_dataframe(df)


@basic.route('/stock_stop', methods=['GET'])
def stock_stop():
    stock_info_sz_delist_df = ak.stock_info_sz_delist(indicator="终止上市公司")
    stock_info_sz_delist_df["status"] = "终止上市"
    stock_info_sz_delist_df_1 = ak.stock_info_sz_delist(indicator="暂停上市公司")
    stock_info_sz_delist_df_1["status"] = "暂停上市"

    df = stock_info_sz_delist_df.append(stock_info_sz_delist_df_1)

    stock_info_sh_delist_df = ak.stock_info_sh_delist(indicator="终止上市公司")
    stock_info_sh_delist_df = stock_info_sh_delist_df[[
        "SECURITY_CODE_A", "SECURITY_ABBR_A", "changeDate"
    ]]
    stock_info_sh_delist_df["status"] = "终止上市"
    stock_info_sh_delist_df1 = ak.stock_info_sh_delist(indicator="暂停上市公司")
    stock_info_sh_delist_df1 = stock_info_sh_delist_df1[[
        "SECURITY_CODE_A", "SECURITY_ABBR_A", "QIANYI_DATE"
    ]]
    df = df.append(stock_info_sh_delist_df1)
    df.columns = ['stockCode', 'companyAbbr', "changeDate", 'status']

    return jsonify_dataframe(df)





if __name__ == '__main__':
    stock_info_sh_df = ak.stock_info_sh_name_code(indicator="主板A股")

    pprint.pprint(stock_info_sh_df)
