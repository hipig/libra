<template>
    <div class="yu-login">
        <div class="yu-login-board">
            <Card :bordered="false" :dis-hover="true">
                <h2 slot="title">Libra - 登陆</h2>
                <Input clearable prefix="ios-contact" placeholder="账号" style="margin-top: 10px;" />
                <Input clearable type="password" prefix="ios-lock" placeholder="密码" style="margin-top: 20px;" />
                <Divider dashed />
                <Button type="primary" long>登陆</Button>
            </Card>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Login"
    }
</script>

<style scoped lang="scss">
    .yu-login{
        background: url("https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/6.jpg");
        background-size: cover;
        width: 100%;
        height: 100%;
        .yu-login-board{
            width: 350px;
            position: absolute;
            top: 40%;
            right: 5%;
            transform: translate3d(-40%,-5%,0);
            opacity: 0.6;
            transition: all 0.5s ease-in;
            &:hover{
                opacity: 1;
            }
        }
    }
</style>
