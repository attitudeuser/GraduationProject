<template>
  <header class="main-header">
    <!-- Logo -->
    <a href="index2.html" class="logo">
      <!-- mini logo for sidebar mini 50x50 pixels -->
      <span class="logo-mini"><b>毕设选题</b></span>
      <!-- logo for regular state and mobile devices -->
      <span class="logo-lg"><b>毕业设计网上选题系统</b></span>
    </a>

    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top">
      <!-- Sidebar toggle button-->
      <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>
      <!-- Navbar Right Menu -->
      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">
          <li class=" user user-menu">
            <button class="text-danger" @click="signout">
              退出登录
            </button>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'va-navibar',
  computed: {
    ...mapGetters([
      'unreadMessagesCount',
      'unreadNotificationsCount',
      'remainTasksCount',
      'currentUser'
    ])
  },
  methods:{
    signout(){
      sessionStorage.removeItem("token");
      sessionStorage.removeItem("uid");
      sessionStorage.removeItem("type");
      sessionStorage.removeItem("username");
      sessionStorage.removeItem("user");
      this.$router.replace({path:'/auth'});
    }
  }
}

</script>

<style>
.user-menu button{
  padding: 15px;
  background: #00a7d0;
  border: none;
  color: #fff;
}
</style>