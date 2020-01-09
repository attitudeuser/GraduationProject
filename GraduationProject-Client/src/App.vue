<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'app',
}
</script>

<style>
.van-tabs__nav{
  background-color: inherit !important;
}

.user_type .van-field__label{
  width: 42px !important;
  font-size: 16px;
  line-height: 39px;
}
.user_type .van-cell{
  padding: 0 !important;
}
.user_type .van-cell__value{
  font-size: 16px;
  line-height: 39px;
  padding-left: 42px;
}
li .van-popup{
  width: 400px;
}
.changform .el-form{
  width: 450px;
  margin: 0 auto;
  margin-top: 37px;
}
.changform .van-popup{
  width: 600px;
  height: 450px;
}
.changform .el-form-item__label{
  width: 120px !important;
}
.changform .el-input{
  width: 300px;
}
.usercourse .label{
  margin-right: 10px;
}
.login .van-popup{
  width: 450px;
  height: 300px;
  text-align: center;
  font-size: 20px;
  line-height: 300px;
}
.el-picker-panel{
  z-index: 2002 !important;
}
</style>
