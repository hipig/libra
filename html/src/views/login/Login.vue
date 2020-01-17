<template>
    <div class="yu-login">
        <div class="yu-login-board">
            <Card :bordered="false" :dis-hover="true">
                <h2 slot="title">Libra - 登陆</h2>
                <Input clearable prefix="ios-contact" maxlength="16" placeholder="账号" v-model="account" style="margin-top: 10px;" />
                <Input clearable type="password" prefix="ios-lock" maxlength="16" v-model="password" placeholder="密码" style="margin-top: 20px;" />
                <Divider dashed />
                <Button type="primary" long @click="onlogin">登陆</Button>
            </Card>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Login",
        data(){
            return {
                account: "admin",
                password:"123456"
            }
        },
        methods:{
            onlogin(){
                if(this.account.trim() == "" || this.password.trim() == ""){
                    this.$Message.warning("账号或密码不能为空!")
                    return
                }
                this.$router.push({
                    name: 'admin'
                })
            }
        }
    }
</script>

<style scoped lang="scss">
    .yu-login{
        /*background: url("https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/6.jpg");*/
        /*background-size: cover;*/
        background: gainsboro;
        width: 100%;
        height: 100%;
        .yu-login-board{
            width: 350px;
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate3d(-50%,-50%,0);
        }
    }
</style>
