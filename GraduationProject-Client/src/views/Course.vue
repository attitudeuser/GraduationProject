<template>
  <div class="usercourse">
    <div>
      <va-todo-list></va-todo-list>
    </div>
  </div>
</template>
<script>
import VATodoList from "../widgets/VATodoList.vue";
export default {
  name: "VAUserCourse",
  data() {
    return {};
  },
  components: {
    "va-todo-list": VATodoList
  },
  created() {
    //验证登陆
    if (!this.IsLogin()) {
      this.$router.push({
        path: this.redirect || "/auth"
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
    }
  }
};
</script>