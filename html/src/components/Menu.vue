<template>
    <div class="libra-menu">
        <BackTop :bottom="140" :right="38">
            <Icon size="40" type="md-arrow-dropup-circle" color="#ffffff" />
        </BackTop>
        <Icon size="36" type="logo-github" color="#ffffff" />
        <Icon @click="value = true" size="36" type="md-settings" color="#ffffff" />
        <Drawer title="设置" placement="left" :closable="false" :scrollable="true" v-model="value">
            <Row>
                主题色：<ColorPicker v-model="theme" alpha />
            </Row>
        </Drawer>
    </div>
</template>

<script>
  export default {
    name: "Menu",
    data () {
      return {
        value: false,
        theme:'#20bf6b',
      }
    },
    methods:{
      trigger(){
        // this.$refs.menu.classList.toggle("active")
      }
    }
  };
</script>

<style scoped lang="scss">
    .libra-menu{
        .ivu-icon-md-arrow-dropup-circle{
            cursor: pointer;
        }
        .ivu-icon-logo-github{
            position: fixed;
            right: 50px;
            bottom: 100px;
            cursor: pointer;
        }
        .ivu-icon-md-settings{
            position: fixed;
            right: 50px;
            bottom: 50px;
            cursor: pointer;
            animation: rotate 2s linear infinite;
        }
    }
    
    @keyframes rotate {
        from{
            transform: rotate(0deg);
        }
        to{
            transform: rotate(360deg);
        }
    }
</style>
