<template>
  <div class="courseitem">
    <div>
      <va-box title="发布课题" theme="box-info" :isBorder="true" widgetType>
        <div slot="content">
          <!-- <form class="form-horizontal">
            <input type="text" id="title" placeholder="请输入题目" class="form-control" v-model="issue.title">
            <va-input
              :isHorizontal="true"
              title="题目"
              placeholder="请输入题目"
              type="text"
              vaId="title"
              :value="issue.title"
            ></va-input>

            <va-input
              :isHorizontal="true"
              title="类型"
              placeholder="请输入课题类型"
              type="text"
              vaId="type"
              :value="issue.type"
            ></va-input>

            <va-input
              :isHorizontal="true"
              title="要求"
              placeholder="请输入课题要求"
              type="text"
              vaId="require"
              :value="issue.require"
            ></va-input>

            <va-textarea-group
              title="内容"
              :isDisabled="false"
              placeholder="请输入课题内容..."
              :value="issue.content"
            >{{issue.content}}</va-textarea-group>
          </form>-->
          <form class="form-horizontal">
            <div class="form-group">
              <label for="title" class="col-sm-2 control-label">题目</label>
              <div class="col-sm-10">
                <input type="text" id="title" placeholder="请输入题目" class="form-control" v-model="issue.title" />
              </div>
            </div>
            <div class="form-group">
              <label for="type" class="col-sm-2 control-label">类型</label>
              <div class="col-sm-10">
                <input type="text" id="type" placeholder="请输入课题类型" class="form-control" v-model="issue.type" />
              </div>
            </div>
            <div class="form-group">
              <label for="require" class="col-sm-2 control-label">要求</label>
              <div class="col-sm-10">
                <input type="text" id="require" placeholder="请输入课题要求" class="form-control" v-model="issue.require" />
              </div>
            </div>
            <div class="form-group" value>
              <label class="col-sm-2 control-label">内容</label>
              <div class="col-sm-10">
                <textarea rows="3" placeholder="请输入课题内容..." class="form-control" v-model="issue.content"></textarea>
              </div>
            </div>
          </form>
        </div>

        <div slot="footer">
          <button type="button" class="btn btn-info pull-right" @click="submit">发布</button>
        </div>
      </va-box>
    </div>
  </div>
</template>
<script>
import VABox from "../widgets/VABox";
import VACheckBox from "../components/VACheckBox";
import VAInput from "../components/VAInput";
import VAButton from "../components/VAButton";
import VATextareaGroup from "../components/VATextareaGroup";
export default {
  name: "VACourseItem",
  data() {
    return {
      issue: {
        title: "",
        type: "",
        require: "",
        cap: 1,
        content: ""
      }
    };
  },
  components: {
    "va-box": VABox,
    "va-checkbox": VACheckBox,
    "va-input": VAInput,
    "va-button": VAButton,
    "va-textarea-group": VATextareaGroup
  },
  created() {
    //验证登陆
    if (!this.IsLogin()) {
      this.$router.push({
        path: this.redirect || "/auth"
      });
    }
    //验证身份
    if (!this.Auth()) {
      alert("只允许教师访问");
      this.$router.push({
        path: this.redirect || "/user/profile"
      });
    }
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
    Auth(){
      if (sessionStorage.getItem("type") !== "teacher") {
        return false;
      }
      return true;
    },
    submit() {
      if (
        this.issue.title === "" ||
        this.issue.type === "" ||
        this.issue.require === "" ||
        this.issue.content === ""
      ) {
        alert("选项不能为空");
        return;
      }
      let param = new URLSearchParams({
        title: this.issue.title,
        type: this.issue.type,
        require: this.issue.require,
        cap: this.issue.cap,
        content: this.issue.content,
        tid: sessionStorage.getItem("uid")
      });
      let _this = this;
      this.$http
        .post(_this.$api+"/api/issue/publish", param)
        .then(function(response) {
          console.log(response.data);
          if (response.data.code === 200) {
            alert("发布成功！");
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