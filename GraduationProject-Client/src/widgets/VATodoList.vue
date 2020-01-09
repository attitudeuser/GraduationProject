<template>
  <!-- TO DO List -->
  <div class="box box-primary">
    <div class="box-header">
      <i class="ion ion-clipboard"></i>
      <h3 class="box-title">选题列表</h3>
    </div>
    <!-- /.box-header -->
    <div class="box-body">
      <ul class="todo-list" ref="table">
        <li v-for="node in list" :key="node.id">
          <!-- <input type="checkbox" class="selbtn" :value="node.id"> -->
          <el-radio :label="node.id" v-model="radio" :disabled="node.cap == 0"></el-radio>
          <span class="text">{{node.title}}</span>
          <small class="label label-info">
            <i class="fa fa-tags"></i>
            容量 {{node.cap}}
          </small>
          要求:&nbsp;&nbsp;
          <span>{{node.require}}</span>
        </li>
      </ul>
    </div>
    <!-- /.box-body -->
    <div class="box-footer clearfix no-border">
      <button type="button" class="btn btn-default pull-right" @click="submit">
        <i class="fa fa-plus"></i> 提交
      </button>
    </div>
  </div>
  <!-- /.box -->
</template>

<script>
export default {
  name: "TodoList",
  data() {
    return {
      list: {},
      checkDom: "",
      radio: 0
    };
  },
  mounted() {
    this.checkDom = document.getElementsByTagName("li");
  },
  created() {
    //获取用户信息
    let _this = this;
    var h = {
      Authorization: "Bearer " + sessionStorage.getItem("token"),
      uid: sessionStorage.getItem("uid")
    };
    this.$http
      .get(_this.$api + `/api/issue/list`, {
        headers: h,
        dataType: "JSONP"
      })
      .then(function(response) {
        _this.list = response.data.data;
      });
  },
  updated() {},
  methods: {
    submit() {
      let _this = this;
      let param = new URLSearchParams({
        iid: this.radio,
        uid: sessionStorage.getItem("uid")
      });
      this.$http
        .post(_this.$api + "/api/subject/select", param)
        .then(function(response) {
          if (response.data.code === 200) {
            alert("选题成功！");
            _this.$router.push({
              path: _this.redirect || "/user/profile"
            });
          }
          return;
        })
        .catch(error => {
          if (error.response.status === 500) {
            alert(error.response.data.message);
          }
        });
    }
  }
};
</script>
<style lang="css" scoped>
.label {
  margin-right: 10px;
}
</style>