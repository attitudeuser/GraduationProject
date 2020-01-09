<template>
  <!-- Chat box -->
  <div class="box box-success">
    <div class="box-header">
      <i class="fa fa-comments-o"></i>

      <h3 class="box-title">交流区</h3>

      <!-- <div class="box-tools pull-right" data-toggle="tooltip" title="Status">
        <div class="btn-group" data-toggle="btn-toggle">
          <button type="button" class="btn btn-default btn-sm active"><i class="fa fa-square text-green"></i>
          </button>
          <button type="button" class="btn btn-default btn-sm"><i class="fa fa-square text-red"></i></button>
        </div>
      </div>-->
    </div>
    <div class="box-body chat id-chat-box">
      <!-- /.item -->
      <!-- chat item -->
      <div class="item" v-for="node in list" :key="node.id">
        <img src="~admin-lte/dist/img/user2-160x160.jpg" alt="user image" class="online" />

        <p class="message">
          <a href="#" class="name">
            <small class="text-muted pull-right">
              <i class="fa fa-clock-o"></i>
              {{node.send_time}}
            </small>
            {{node.username}}
          </a>
          {{node.message}}
        </p>
      </div>
      <!-- /.item -->
    </div>
    <!-- /.chat -->
    <div class="box-footer">
      <div class="input-group">
        <input class="form-control" placeholder="请输入您要发送的内容..." v-model="content" />

        <div class="input-group-btn">
          <button type="button" class="btn btn-success" @click="Send">发送</button>
        </div>
      </div>
    </div>
  </div>
  <!-- /.box (chat box) -->
</template>

<script>
export default {
  name: "ChatBox",
  data() {
    return {
      list: null,
      content: ""
    };
  },
  mounted() {
    var chatBoxElement = this.$el.querySelector(".id-chat-box");
    $(chatBoxElement).slimScroll({
      height: "500px"
    });
  },
  created() {
    //验证登陆
    if (!this.IsLogin()) {
      this.$router.push({
        path: this.redirect || "/auth"
      });
    }
    //获得消息
    this.GetMessage();
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
    GetMessage() {
      let _this = this;
      this.$http
        .get(_this.$api+"/api/message/list")
        .then(function(response) {
          if (response.data.code === 200) {
            _this.list = response.data.data;
            console.log(_this.message);
          }
          return;
        })
        .catch(error => {
          console.log(error.response);
          return;
        });
    },
    Send() {
      let _this = this;
      let param = new URLSearchParams({
        message: this.content,
        fid: sessionStorage.getItem("uid"),
        username: sessionStorage.getItem("username")
      });
      this.$http
        .post(_this.$api+"/api/message/send", param)
        .then(function(response) {
          if (response.data.code === 200) {
            // alert("发布成功！");
            _this.GetMessage();
            _this.content = '';
            _this.$router.push({
              path: _this.redirect || "/user/communication"
            });
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
.box-success {
  height: 600px;
}
.box-body {
  height: 500px !important;
}
</style>
