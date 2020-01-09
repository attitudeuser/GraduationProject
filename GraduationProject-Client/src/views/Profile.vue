<template>
  <div class="userinfor">
    <div class="userinforation">
      <div class="col-md-4">
        <va-social-user-v1
          name="个人信息"
          :description="user.description"
          profileImage="/static/img/user1-128x128.jpg"
          :infoList="socialDataList[0]"
        ></va-social-user-v1>
        <div class="changeinfo">
          <button @click="showpopup" class="btn btn-success">修改个人信息</button>
        </div>
      </div>
    </div>
    <div class="changform">
      <van-popup v-model="show">
        <el-form :model="ruleForm" ref="ruleForm" label-width="100px" class="demo-ruleForm">
          <!-- <el-form-item label="姓名" prop="name">
            <el-input disabled :value="user.username"></el-input>
          </el-form-item>-->
          <el-form-item label="学号/教职工号" prop="name">
            <el-input v-model="ruleForm.sid" placeholder="请输入学号/教职工号"></el-input>
          </el-form-item>
          <el-form-item label="年龄" prop="age">
            <el-input v-model="ruleForm.age" placeholder="请输入年龄"></el-input>
          </el-form-item>
          <el-form-item label="学院" prop="region">
            <el-input v-model="ruleForm.college" placeholder="请输入所在学院完整名称"></el-input>
          </el-form-item>
          <el-form-item label="入学/入职时间" prop="grade">
            <el-input v-model="ruleForm.grade" placeholder="选择入学/入职年份"></el-input>
          </el-form-item>
          <el-form-item label="联系方式" prop="phone">
            <el-input v-model="ruleForm.phone" placeholder="请输入手机号"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm">保存修改</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </van-popup>
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
            <div v-if="subject_select === true">
              <va-callout type="info" :title="issue.issue_title" :text="issue.issue_text"></va-callout>
            </div>
            <div v-else>
              <va-callout type="danger" title="尚未选择题目" text="您还没有选择毕设题目！"></va-callout>
            </div>
          </div>
          <!-- /.box-body -->
        </div>
        <!-- /.box -->
      </div>
    </div>
  </div>
</template>
<script>
import VASocialUserV1 from "../widgets/VASocialUser.v1.vue";
import VACallout from "../components/VACallout.vue";
export default {
  name: "VAUserInfor",
  data() {
    return {
      user: sessionStorage.getItem("user"),
      show: false,
      user: {},
      subject: {},
      issue: {},
      subject_select: false,
      socialDataList: [
        [
          {
            count: "姓名",
            text: ""
          },
          {
            count: "学号",
            text: ""
          },
          {
            count: "学院",
            text: ""
          }
        ]
      ],
      ruleForm: {
        sid: "",
        college: "",
        grade: "",
        phone: "",
        age: 0
      },
      rules: {
        name: [
          { required: true, message: "请填写学号/员工号", trigger: "blur" },
          { min: 8, max: 11, message: "长度在 8 到 11 个字符", trigger: "blur" }
        ],
        region: [
          {
            required: true,
            message: "请填写所在学院的完整名称",
            trigger: "blur"
          }
        ],
        phone: [
          { required: true, message: "请填写手机号", trigger: "blur" },
          {
            min: 11,
            max: 11,
            message: "请检查手机号的长度是否正确",
            trigger: "blur"
          }
        ],
        date1: [
          {
            type: "date",
            required: true,
            message: "请选择年份",
            trigger: "change"
          }
        ]
      }
    };
  },
  components: {
    "va-social-user-v1": VASocialUserV1,
    "va-callout": VACallout
  },
  created() {
    //验证登陆

    if (!this.IsLogin()) {
      this.$router.push({
        path: this.redirect || "/auth"
      });
    }
    //获取用户信息
    this.GetProfile();
    //获取选课信息
    this.GetSubject();
  },
  methods: {
    IsLogin() {
      if (
        sessionStorage.getItem("token") === null ||
        sessionStorage.getItem("uid") === null
      ) {
        return false;
      }
      return true;
    },
    //获得个人信息
    GetProfile() {
      //获取用户信息
      let _this = this;
      var h = {
        Authorization: "Bearer " + sessionStorage.getItem("token"),
        uid: sessionStorage.getItem("uid")
      };
      this.$http
        .get(_this.$api + `/api/user/profile`, {
          headers: h,
          dataType: "JSONP"
        })
        .then(function(response) {
          _this.user = response.data.data;
          _this.socialDataList[0][0].text = _this.user.username;
          _this.socialDataList[0][1].text = _this.user.sid;
          _this.socialDataList[0][2].text = _this.user.college;
          if (_this.user.type === "admin") {
            _this.user.type = "管理员";
          } else if (_this.user.type === "student") {
            _this.user.type = "学生";
          } else if (_this.user.type === "teacher") {
            _this.user.type = "教师";
          }
          _this.user.description = _this.user.grade + " & " + _this.user.type;
        })
        .catch(error => {
          if (error.response.status === 400) {
            alert("token失效");
            _this.$router.push({
              path: _this.redirect || "/auth"
            });
          }
          return;
        });
    },
    //获得选课信息
    GetSubject() {
      //获取用户信息
      let _this = this;
      var h = {
        Authorization: "Bearer " + sessionStorage.getItem("token"),
        uid: sessionStorage.getItem("uid")
      };
      this.$http
        .get(
          _this.$api + "/api/subject/get?uid=" + sessionStorage.getItem("uid"),
          {
            headers: h,
            dataType: "JSONP"
          }
        )
        .then(function(response) {
          _this.subject = response.data.data;
          if (_this.subject.id !== 0) {
            _this.subject_select = true;
            _this.GetIssue(_this.subject.iid);
            sessionStorage.setItem("iid", _this.subject.iid);
            sessionStorage.setItem("sid", _this.subject.id);
          } else {
            _this.subject_select = false;
          }
        });
    },
    //表单
    submitForm() {
      let param = new URLSearchParams({
        sid: this.ruleForm.sid,
        college: this.ruleForm.college,
        grade: this.ruleForm.grade,
        phone: this.ruleForm.phone,
        age: this.ruleForm.age,
        id: sessionStorage.getItem("uid")
      });
      var h = {
        Authorization: "Bearer " + sessionStorage.getItem("token"),
        uid: sessionStorage.getItem("uid")
      };
      let _this = this;
      this.$http
        .post(_this.$api + "/api/user/update", param, {
          headers: h,
          dataType: "JSONP"
        })
        .then(function(response) {
          if (response.data.code === 200) {
            alert("修改成功！");
          }
          _this.show = false;
          /*_this.$router.push({
            path: _this.redirect || "/user/profile"
          });*/
          _this.$router.go(0);
          return;
        })
        .catch(error => {
          console.log(error.response);
          return;
        });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    //展示表单弹框
    showpopup() {
      this.show = true;
    },
    async GetIssue(id) {
      //获取用户信息
      var _this = this;
      var h = {
        Authorization: "Bearer " + sessionStorage.getItem("token"),
        uid: sessionStorage.getItem("uid")
      };
      this.$http
        .get(_this.$api + "/api/subject/relate?iid=" + id, {
          headers: h,
          dataType: "JSONP"
        })
        .then(function(response) {
          _this.issue = response.data.data;

          _this.GetUser(_this.issue.tid).then(res => {
            sessionStorage.setItem("teacher_name", res.username);
            sessionStorage.setItem("teacher_phone", res.phone);
          });
          _this.issue.issue_title = "成功选择 " + _this.issue.title + " 题目";
          _this.issue.issue_text =
            " 题目 指导老师: " +

            sessionStorage.getItem("teacher_name") +
            " 联系方式: " +
            sessionStorage.getItem("teacher_phone");
        });
    },
    GetUser(id) {
      var that = this;
      var h = {
        uid: id
      };
      return new Promise((resolve, reject) => {
        that.$http
          .get(that.$api + "/api/general/profile", {
            headers: h
          })
          .then(res => {
            resolve(res.data.data);
          });
      });
    }
  }
};
</script>
<style lang="css" scoped>
.clearfix:after {
  content: "";
  display: block;
  clear: both;
}
.clearfix {
  zoom: 1;
}
.col-md-4 {
  width: 600px;
  position: relative;
}
.col-md-4 .changeinfo {
  position: absolute;
  top: 16px;
  right: 24px;
}
.col-md-4 .changeinfo button {
  width: 130px;
  height: 40px;
  background-color: #0096d4 !important;
  border-color: #0096d4 !important;
}
.userinforation {
  padding-top: 50px;
}
</style>