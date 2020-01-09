<template>
  <div class="login">
    <div class="login_form">
      <div class="title">
        <p>毕业设计网上选题系统</p>
      </div>
      <van-tabs>
        <van-tab title="已有账号，立即登录">
          <form>
            <ul>
              <li>
                <div class="user_name">
                  <img src="../images/username.png" />
                  <input type="text" name placeholder="请输入用户名" v-model="user.username" />
                </div>
              </li>
              <li>
                <div class="user_password">
                  <img src="../images/userpsw.png" />
                  <input type="password" name placeholder="请输入密码" v-model="user.password" />
                </div>
              </li>
              <!-- <li>
                <div class="forget_psw">
                  <p @click="showway">忘记密码</p>
                </div>
                <van-popup v-model="forpsw">
                  <h3>通过绑定的邮箱来找回密码</h3>
              
                  
                </van-popup>
              </li> -->
            </ul>
            <div class="submit_go">
              <button type="button" @click="signin">立即登录</button>
            </div>
          </form>
        </van-tab>
        <van-tab title="新用户注册">
          <form>
            <ul>
              <li>
                <div class="user_name">
                  <img src="../images/username.png" />
                  <input type="text" name placeholder="请输入姓名" required v-model="register.username" />
                </div>
              </li>
              <li>
                <div class="user_name">
                  <img src="../images/email.png" />
                  <input type="text" name placeholder="请输入邮箱" required v-model="register.email" />
                </div>
              </li>
              <li>
                <div class="user_password">
                  <img src="../images/userpsw.png" />
                  <input type="password" name placeholder="请输入密码" required v-model="register.password" />
                </div>
              </li>
              <li>
                <div class="user_password">
                  <img src="../images/userpsw.png" />
                  <input type="password" name placeholder="重复输入密码" required v-model="register.repassword" />
                </div>
              </li>
              <li>
                <div class="user_type">
                  <img src="../images/usertype.png" />
                  <van-field
                    readonly
                    clickable
                    :value="type"
                    placeholder="绑定登录身份"
                    @click="showPicker = true"
                  ></van-field>
                </div>
                <van-popup v-model="showPicker">
                  <van-picker
                    show-toolbar
                    :columns="columns"
                    @cancel="showPicker = false"
                    @confirm="onConfirm"
                  ></van-picker>
                </van-popup>
              </li>
            </ul>
            <div class="submit_go">
              <button type="button" @click="signup">立即注册</button>
            </div>
          </form>
        </van-tab>
        <van-tab title="找回密码">
            <form action="">
              <ul>
                <li>
                  <div class="user_name">
                    <img src="../images/email.png" />
                    <input type="text" name placeholder="请输入邮箱" required v-model="register.email" />
                  </div>
                </li>
              </ul>
              <div class="submit_go">
                <button type="button" @click="send">立即找回</button>
              </div>
            </form>

        </van-tab>

      </van-tabs>
    </div>
  </div>
</template>
<script>
export default {
  name: "login",
  data() {
    return {
      forpsw: false,
      user: {
        username: "",
        password: ""
      },
      register: {
        username: "",
        password: "",
        repassword: "",
        email: ""
      },
      type: "",
      columns: ["学生", "教师"],
      showPicker: false,
      showPicker1: false
    };
  },
  methods: {
    onConfirm(value) {
      this.type = value;
      this.showPicker = false;
    },
    onConfirm1(value) {
      this.type = value;
      this.showPicker1 = false;
    },
    showway() {
      this.forpsw = true;
    },
    signup() {
      if (
        this.register.username === "" ||
        this.register.password === "" ||
        this.register.repassword === "" ||
        this.register.email === ""
      ) {
        alert("选项不能为空");
        return;
      }
      if (this.register.password != this.register.repassword) {
        alert("前后两次密码不正确");
        return;
      }
      //  检测完成
      if (this.type === "学生") {
        this.type = "student"
      }else if(this.type === "教师"){
        this.type = "teacher"
      }
      let param = new URLSearchParams({
        username: this.register.username,
        password: this.register.password,
        email: this.register.email,
        type: this.type
      });
      let _this = this;
      this.$http
        .post(_this.$api+"/api/general/signup", param)
        .then(function(response) {
          console.log(response.data);
          if (response.data.code === 200) {
            alert("注册成功！")
          }
          return;
        })
        .catch(error => {
          console.log(error.response);
          return;
        });
    },
    signin() {
      if (this.user.username === "" || this.user.password === "") {
        alert("选项不能为空");
        return;
      }
      //开始登陆请求
      //开始登陆逻辑 将token记录 存入sessionStorage中
      let param = new URLSearchParams({
        username: this.user.username,
        password: this.user.password
      });
      let _this = this;
      this.$http
        .post(_this.$api+"/api/general/signin", param)
        .then(function(response) {
          console.log(response.data);
          if (response.data.code === 200) {
            sessionStorage.setItem("token", response.data.message);
            sessionStorage.setItem("uid", response.data.data.id);
            sessionStorage.setItem("type", response.data.data.type);
            sessionStorage.setItem("username", response.data.data.username);
            sessionStorage.setItem("user", response.data.data);
            //跳转到后台
            _this.$router.push({
              path: _this.redirect || "/user/profile"
            });
          }
          return;
        })
        .catch(error => {
          console.log(error.response);
          return;
        });
    },
    send(){
      let _this = this;
      this.$http
        .get(_this.$api+"/api/general/forget?email="+_this.register.email)
        .then(function(response) {
          console.log(response.data);
          if (response.data.code === 200) {
            /*sessionStorage.setItem("token", response.data.message);
            sessionStorage.setItem("uid", response.data.data.id);
            sessionStorage.setItem("type", response.data.data.type);
            sessionStorage.setItem("username", response.data.data.username);
            sessionStorage.setItem("user", response.data.data);*/
            alert("邮件发送成功，请注意查收。");
            //跳转到后台
            /*_this.$router.push({
              path: _this.redirect || "/user/profile"
            });*/
          }
          return;
        })
        .catch(error => {
          console.log(error.response);
          return;
        });
    }
  }
};
</script>
<style lang="css" scoped>
.login {
  width: 100%;
  height: 100%;
  /*background-image: url(../../images/bg1.jpg);*/
  background-repeat: no-repeat;
  background-size: cover;
  position: relative;
}
.login_form {
  width: 500px;
  height: 500px;
  margin: 0 auto;
  padding-top: 100px;
}
.title {
  padding: 10px;
  margin-bottom: 20px;
  text-align: center;
}
.title p {
  font-size: 30px;
  font-weight: bolder;
}
form {
  margin-top: 40px;
}
form ul li {
  width: 100%;
  height: 60px;
}
.user_name,
.user_password, .user_email{
  width: 400px;
  margin: 0 auto;
  height: 40px;
  position: relative;
}
.user_name img,
.user_password img, .user_email img{
  width: 20px;
  height: 20px;
  position: absolute;
  top: 50%;
  left: 11px;
  transform: translate(0, -50%);
}
.user_name input,
.user_password input,.user_email input{
  width: 100%;
  height: 100%;
  border: 1px solid #eee;
  font-size: 16px;
  padding-left: 42px;
}
.user_email input{
  padding-left: 10px;
  display: block;
}
.user_password p {
  color: red;
  padding-left: 10px;
}
.forget_psw {
  width: 100%;
}
.forget_psw p {
  float: right;
}
.forget_psw p:hover {
  cursor: pointer;
}
.user_type {
  position: relative;
  width: 400px;
  height: 40px;
  margin: 0 auto;
  border: 1px solid #eee;
}
.user_type img {
  width: 20px;
  height: 20px;
  position: absolute;
  top: 50%;
  left: 11px;
  transform: translate(0, -50%);
  z-index: 10;
}
.submit_go {
  width: 400px;
  margin: 0 auto;
  height: 50px;
}
.submit_go button {
  width: 100%;
  height: 100%;
  background-color: inherit;
  border: 2px solid #000;
  border-radius: 20px;
  font-size: 16px;
  letter-spacing: 1px;
}
.submit_go button:hover {
  cursor: pointer;
}


</style>