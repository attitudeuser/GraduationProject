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
									<img src="../images/username.png">
									<input type="text" name="" placeholder="请输入用户名">
								</div>
							</li>
							<li>
								<div class="user_password">
									<img src="../images/userpsw.png">
									<input type="password" name="" placeholder="请输入密码">
								</div>
							</li>
							<li>
								<div class="user_type">
									<img src="../images/usertype.png">
									<van-field readonly clickable :value="value" placeholder="请选择登录身份" @click="showPicker1 = true">
										
									</van-field>
								</div>
								<van-popup v-model="showPicker1">
									<van-picker show-toolbar :columns="columns" @cancel="showPicke1 = false" @confirm="onConfirm1"></van-picker>
								</van-popup>
							</li>
							<li>
								<div class="forget_psw">
									<p @click="showway">忘记密码</p>
								</div>
							</li>
						</ul>
						<div class="submit_go">
							<button>立即登录</button>
						</div>
					</form>
				</van-tab>
				<van-tab title="新用户注册">
					<form>
						<ul>
							<li>
								<div class="user_name">
									<img src="../images/username.png">
									<input type="text" name="" placeholder="请输入学号/教职工号">
								</div>
							</li>
							<li>
								<div class="user_password">
									<img src="../images/userpsw.png">
									<input type="password" name="" placeholder="请输入密码">
									<p>两次输入密码不一致！</p>
								</div>
							</li>
							<li>
								<div class="user_password">
									<img src="../images/userpsw.png">
									<input type="password" name="" placeholder="重复输入密码">
								</div>

							</li>
							<li>
								<div class="user_type">
									<img src="../images/usertype.png">
									<van-field readonly clickable :value="value" placeholder="绑定登录身份" @click="showPicker = true">
										
									</van-field>
								</div>
								<van-popup v-model="showPicker">
									<van-picker show-toolbar :columns="columns" @cancel="showPicker = false" @confirm="onConfirm"></van-picker>
								</van-popup>
							</li>
						</ul>
						<div class="submit_go">
							<button>立即注册</button>
						</div>
					</form>
				</van-tab>

			</van-tabs>
			
		</div>
	</div>
</template>
<script>
	export default{
		name: 'login',
		data(){
			return{
				value: '',
				columns: ['学生','教师'],
				showPicker: false,
				showPicker1: false,
			}
		},
		methods:{
			onConfirm(value){
				this.value = value;
				this.showPicker = false;
			},
			onConfirm1(value){
				this.value = value;
				this.showPicker1 = false;
			},
			showway(){
			}
		}
	}
</script>
<style lang="css" scoped>
	.login{
		width: 100%;
		height: 100%;
		/*background-image: url(../../images/bg1.jpg);*/
		background-repeat: no-repeat;
		background-size: cover;
		position: relative;
	}
	.login_form{
		width: 500px;
		height: 500px;
		margin: 0 auto;
		padding-top: 100px;
	}
	.title{
		padding: 10px;
		margin-bottom: 20px;
		text-align: center;
		
	}
	.title p{
		font-size: 30px;
		font-weight: bolder;
	}
	form{
		margin-top: 40px;
	}
	form ul li{
		width: 100%;
		height: 60px;
	}
	.user_name,.user_password{
		width: 400px;
		margin: 0 auto;
		height: 40px;
		position: relative;
		
	}
	.user_name img,.user_password img{
		width: 20px;
		height: 20px;
		position: absolute;
		top: 50%;
		left: 11px;
		transform: translate(0,-50%);
	}
	.user_name input,.user_password input{
		width: 100%;
		height: 100%;
		border: 1px solid #eee;
		font-size: 16px;
		padding-left: 42px;
	}
	.user_password p{				
		color: red;
		padding-left: 10px;
	}
	.forget_psw{
		width: 100%;
		
	}
	.forget_psw p{
		float: right;
	}
	.forget_psw p:hover{
		cursor: pointer;
	}
	.user_type{
		position: relative;
		width: 400px;
		height: 40px;
		margin: 0 auto;
		border: 1px solid #eee;
		
	}
	.user_type img{
		width: 20px;
		height: 20px;
		position: absolute;
		top: 50%;
		left: 11px;
		transform: translate(0,-50%);
		z-index: 10;

	}
	.submit_go{
		width: 400px;
		margin: 0 auto;
		height: 50px;
		
	}
	.submit_go button{
		width: 100%;
		height: 100%;
		background-color: inherit;
		border: 2px solid #fff;
		border-radius: 20px;
		font-size: 16px;
		letter-spacing: 1px;
	}
	.submit_go button:hover{
		cursor: pointer;
	}
</style>