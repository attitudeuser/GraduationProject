<template>
	<div class="reset">
		<div class="resetinput">
			<h3>修改密码</h3>
			<div class="user_password">
              <img src="../images/userpsw.png" />
              <input type="password" name placeholder="请输入新密码" required v-model="register.password" />
            </div>
            <div class="user_password">
              <img src="../images/userpsw.png" />
              <input type="password" name placeholder="请重复密码" required v-model="register.repassword" />
            </div>
		</div>
		<div class="submit_go">
	        <button type="button" @click="signup">确定</button>
	    </div>

	</div>
</template>
<script>
	export default{
		name: 'Reset',
		data(){
			return{
				register: {
		        password: "",
		        repassword: ""
		      },
			}
		},
		created(){
			console.log(this.$route.query);
		},
		methods:{
			signup(){
				 if (
			        this.register.password === "" ||
			        this.register.repassword === ""
			      ) {
			        alert("选项不能为空");
			        return;
			      }
			      if (this.register.password != this.register.repassword) {
			        alert("前后两次密码不正确");
			        return;
			      }
			      let _this = this;
			      var h = {
			        Authorization: "Bearer " + sessionStorage.getItem("token"),
			        uid: sessionStorage.getItem("uid")
			      };
			      this.$http
			        .get(_this.$api+"/api/general/reset?password="+_this.register.password+"&uuid="+_this.$route.query.uuid,{
			        	headers: h,
         			 	dataType: "JSONP"
			        })
			        .then(function(response) {
			          console.log(response.data);
			          if (response.data.code === 200) {
			            /*sessionStorage.setItem("token", response.data.message);
			            sessionStorage.setItem("uid", response.data.data.id);
			            sessionStorage.setItem("type", response.data.data.type);
			            sessionStorage.setItem("username", response.data.data.username);
			            sessionStorage.setItem("user", response.data.data);*/
			            alert("密码修改成功！");

			            //跳转到后台
			            _this.$router.push({
			              path: _this.redirect || "/auth"
			            });
			          }
			          return;
			        })
			        .catch(error => {
			          console.log(error.response);
			          alert(error.response.data.message);
			          return;
			        });
			}
		}
	}
</script>
<style lang="css" scoped>
	.reset{
		width: 400px;
		height: 400px;
		margin: 0 auto;
		margin-top: 130px;
	}
	.user_password, .user_email{
	  width: 400px;
	  margin: 0 auto;
	  height: 40px;
	  position: relative;
	  margin-top: 20px;
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
	.submit_go {
	  width: 400px;
	  margin: 0 auto;
	  height: 50px;
	  margin-top: 30px;
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