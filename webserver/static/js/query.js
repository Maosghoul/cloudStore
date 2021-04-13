var httpAddr = "http://121.5.245.69:6789/cs/";
var LoginAddr = "http://121.5.245.69:6789/login";

var arr = [];
var checkBoxIdList = new Array();

function checkUser() {
    var username = $('#uname').val()
    var pwd = $('#pwd').val()

    var reqUrl = httpAddr + "login"
    var cmd = {
        'username': username,
        'password': pwd
    }
    let res = util.httpPost(reqUrl, JSON.stringify(cmd))
    console.log(res)
    if (res == true) {
        util.setItem('username', username)
        return true
    }
    alert("用户名或密码错误")
    return false
}

function regInfo() {
    var username = $('#uname').val()
    var pwd = $('#pwd').val()
    var pwdr = $('#pwdr').val()

    console.log("username:", username)
    console.log("pwd:", pwd)
    console.log("pwdr:", pwdr)

    if (pwd != pwdr) {
        alert("密码与确认密码不一致")
        return false
    }

    var reqUrl = httpAddr + "register"
    var cmd = {
        'username': username,
        'password': pwd,
    }
    let res = util.httpPost(reqUrl, JSON.stringify(cmd))
    console.log(res)
    if (res == true) {
        alert("注册成功")
        return true
    }
    alert("请求错误，用户名已注册")
    return false
}

function inputFile() {
    var data = new FormData();
    username = util.getItem('username')
    filename = $('#upfile').get(0).files[0]
    data.append('username', username)
    data.append('file', filename)
    console.log("username:", username, "filename:", filename)
    console.log("formdata:", data)
    reqUrl = httpAddr + 'upload_file'
    let result = util.httpFormData(reqUrl, data)
    if (result == true) {
        alert("上传成功")
        window.location.replace("http://121.5.245.69:6789/index")

    } else {
        alert("上传失败")
    }
}

function ListFile() {
    var username = util.getItem('username')
    console.log("username:", username)
    var reqUrl = httpAddr + "list_file"
    var cmd = {
        'username': username
    }

    let res = util.httpPostWithResp(reqUrl, JSON.stringify(cmd))
    console.log("res", res.fileinfo)
    if (res == null) {
        alert("网络错误")
    }
    // for (var i = 0; i < res.fileinfo.length; i++) {
    //     console.log("res:", res.fileinfo[i].name, res.fileinfo[i].time)
    // }
    for (var i = 0; i < res.fileinfo.length; i++) {
        arr.push({name: res.fileinfo[i].name, time: res.fileinfo[i].time})
    }

}

function checkLogin() {
    var username = util.getItem("username")
    if (username == null) {
        window.location.href = LoginAddr
        alert("请先登录")
    }
}

function cancel() {
    util.removeItem("username")
    window.location.href = LoginAddr
}

function deleteFile() {
    var username = util.getItem("username")
    if (username == null) {
        window.location.href = LoginAddr
        alert("请先登录")
    }
    var filename = []
    for (var i = 0; i < arr.length; i++) {
        if (checkBoxIdList[i + 1] == 1) {
            filename.push(arr[i].name)
        }
    }
    var reqUrl = httpAddr + "delete_file"
    var cmd = {
        'username': username,
        'filename': filename
    }
    let res = util.httpPost(reqUrl, JSON.stringify(cmd))
    console.log("delete file:", res)
    if (res == true) {
        alert("删除成功")
        window.location.replace("http://121.5.245.69:6789/index")
        return true
    }
    alert("删除失败")
    return false
}

function adultFile() {
    var username = util.getItem("username")
    if (username == null) {
        window.location.href = LoginAddr
        alert("请先登录")
    }
    var filename = []
    for (var i = 0; i < arr.length; i++) {
        if (checkBoxIdList[i + 1] == 1) {
            filename.push(arr[i].name)
        }
    }
    var reqUrl = httpAddr + "adult_file"
    var cmd = {
        'username': username,
        'filename': filename
    }
    let res = util.httpPostWithResp(reqUrl, JSON.stringify(cmd))
    console.log("delete file:", res)
    console.log("res code :",res.code)
    console.log("res message :",res.message)
    if (res.code =="0"){
        alert("所选文件通过审计")
    }else{
        alert("所选文件"+res.message+"未通过审计")
    }
    window.location.replace("http://121.5.245.69:6789/index")
}
