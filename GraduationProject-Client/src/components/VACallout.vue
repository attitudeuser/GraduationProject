<template>
  <div class="callout" :class="classObject">
    <h4>{{ title }}</h4>

    <p>{{ text }}</p>
    <div class="input-group-btn" v-if="this.type === 'info'">
      <button type="button" class="btn btn-success" @click="submit">取消选择此题目</button>
    </div>
  </div>
</template>

<script>
const CalloutTypes = ["success", "info", "warning", "danger"];
export default {
  name: "va-callout",
  props: {
    type: {
      default: "success",
      validator: function(value) {
        return CalloutTypes.includes(value);
      }
    },
    title: {
      type: String
    },
    text: {
      type: String
    }
  },
  computed: {
    classObject() {
      return {
        "callout-success": this.type === "success",
        "callout-info": this.type === "info",
        "callout-warning": this.type === "warning",
        "callout-danger": this.type === "danger"
      };
    }
  },
  methods: {
    submit() {
      var _this = this;
      var sid = sessionStorage.getItem("sid");
      var iid = sessionStorage.getItem("iid");
      this.$http
        .get(_this.$api + "/api/subject/delete?sid=" + sid + "&iid=" + iid)
        .then(function(response) {
          if (response.data.code === 200) {
            alert("取消成功!");
            /*_this.$router.push({
              path: _this.redirect || "/user/profile"
            });*/
            _this.$router.go(0);
          }else{
            alert("操作失败!");
            console.log(response.data);
            return;
          }
        });
    }
  }
};
</script>
<style lang="css" scoped>
button {
  background-color: #0096d4 !important;
  border-color: #0096d4 !important;
}
</style>
