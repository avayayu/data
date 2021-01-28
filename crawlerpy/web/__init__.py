'''
'''
from flask import Flask, jsonify, request
from flask_cors import CORS
from web.basic import basic
app = Flask(__name__)

# flask的跨域解决
# 解读：resources：全局配置允许跨域的API接口
# Cors需要在后端应用进行配置，因此，是一种跨域的后端处理方式
# 此句的理解：前端的访问可能会出现跨域的情况，那么，此句就是为了解决这个问题，
# 一个不认识的源来访问服务端应用时，应用进行授权
CORS(app)
app.register_blueprint(basic)
