<template>
	<div class="userinfor">
		<div class="userinforation">
			<div class="col-md-4">
		        <va-social-user-v1
		          name="个人信息"
		          description="2016级 &amp; 学生"
		          profileImage="/static/img/user1-128x128.jpg"
		          :infoList="socialDataList[0]"
		        ></va-social-user-v1>
		    </div>
		</div>
		<div class="usercourse">
			<div class="col-md-6">
		        <div class="box box-default">
		          <div class="box-header with-border">
		            <i class="fa fa-bullhorn"></i>

		            <h3 class="box-title">选题结果</h3>
		          </div>
		          <!-- /.box-header -->
		          <div class="box-body">

		            <va-callout
		              type="danger"
		              title="尚未选择题目"
		              text="您还没有选择毕设题目！"
		            ></va-callout>

		            <va-callout
		              type="info"
		              title="成功选择某某某某某某某某某某某题目"
		              text="某某某某某某某某某某某题目 指导老师 某某某 联系方式11111111111"
		            ></va-callout>
		          </div>
		          <!-- /.box-body -->
		        </div>
		        <!-- /.box -->
		    </div>
		</div>
	</div>
</template>
<script>
	import VASocialUserV1 from '../widgets/VASocialUser.v1.vue'
	import VACallout from '../components/VACallout.vue'
	export default{
		name: 'VAUserInfor',
		data(){
			return{
				socialDataList: [
			        [
			          {
			            count: '姓名',
			            text: '张三'
			          },
			          {
			            count: '学号',
			            text: '20161101001'
			          },
			          {
			            count: '学院',
			            text: '网络空间安全与计算机学院'
			          },
			          {
			            count: '专业',
			            text: '信息安全'
			          }
			        ]
			    ]
			}
		},
		components: {
		    'va-social-user-v1': VASocialUserV1,
		    'va-callout': VACallout
		},
		created(){

		},
		methods:{
			initData(){

			}
		}
	}
</script>
<style lang="css" scoped>
	.clearfix:after {
	  content: '';
	  display: block;
	  clear: both;
	}
	.clearfix {
	  zoom: 1;
	}
	.col-md-4{
		width: 600px;
	}
	 .userinforation{
		padding-top: 50px;
	}
</style>