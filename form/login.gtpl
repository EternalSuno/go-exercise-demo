<html>
<head>
    <title>登录</title>
</head>
<body>
<form action="/login" enctype="multipart/form-data" method="post">
   <p>用户名:<input type="text" name="username"/></p>
   <p>密码:<input type="password" name="password"/></p>
   <p>真实姓名:<input type="text" name="realname"/></p>
   <p>英文姓名:<input type="text" name="engname"/></p>
   <p>邮箱:<input type="text" name="email"/></p>
   <p>手机号:<input type="text" name="mobile"/></p>
   <p>年龄:<input type="text" name="age"/></p>
   <p>
      <select name="fruit">
         <option value="apple">apple</option>
         <option value="pear">pear</option>
         <option value="banana">banana</option>
      </select>
   </p>
   <p>
      <input type="radio" name="gender" value="1"/>男<br/>
      <input type="radio" name="gender" value="2"/>女<br/>
   </p>
   <input type="checkbox" name="interest" value="football"/> 足球
   <input type="checkbox" name="interest" value="basketball"/> 篮球
   <input type="checkbox" name="interest" value="tennis"/> 网球
   <input type="hidden" name="token" value="{{token}}" />
   <br/>
   <input type="file" name="uploadfile"/>


   <p><input type="submit" value="登录"/></p>
</form>
</body>
</html>