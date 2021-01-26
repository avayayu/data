from flask import jsonify, Response


def jsonify_success(param=None):
    success = {"status": "success"}
    if param:
        success.update(param)
    return jsonify(success)


def jsonify_failure(param):
    return jsonify({"status": "failure", "info": param})


def jsonify_dataframe(df=None):
    return Response(df.to_json(orient="records"), mimetype='application/json')
