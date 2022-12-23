<template>
    <div class="welcome-page">
        <img class="welcome-bg" src="/images/welcome-bg.jpg">
        <div class="tagline-box">
            <div class="logo">
                <span style="color:#434343">NeSU.</span><span style="color:#87C644">Еду</span><span style="color:#9ED5F0">*</span>
            </div>
            <div class="tagline">
                <span style="color:#9ED5F0">*</span><span style="color:#434343"> Мы не сделаем вас сытыми</span><br>
                <span style="color:#434343">&nbsp;&nbsp;&nbsp;Мы научим вас заказывать </span><span style="color:#87C644">еду</span>
            </div>
        </div>
        <div class="log-reg-btns">
            <span>Личный кабинет</span><br>
            <div style="margin: auto; 
                        display: flex; 
                        justify-content:space-between; 
                        margin-top: 13.121%;">
                <div style="display: flexbox; justify-content: flex-start; text-align: center;">
                    <va-button 
                    class="log-reg-btn"
                    color="customGreen"
                    @click="showAuth"
                    >
                        Sign In
                    </va-button><br>
                    <span style="font-size: 35%; font-weight: lighter;">Я уже смешарик</span>
                </div>
                <div style="justify-content: flex-end; text-align: center;">
                    <va-button 
                    class="log-reg-btn"
                    color="customGreen"
                    @click="showSignUp"
                    >
                        Sign Up
                    </va-button><br>
                    <span style="font-size: 35%; font-weight: lighter;">Я новенький</span>
                </div>
            </div>
        </div>
        <div class="guide">
            <div class="guide-text">
                <span>Проголодался&nbsp;&nbsp;&nbsp;</span><img style="width: 78px; height: 25px;" src="/images/right-arrow.svg">
                <span>&nbsp;&nbsp;&nbsp;Заказал&nbsp;&nbsp;&nbsp;</span><img style="width: 78px; height: 25px;" src="/images/right-arrow.svg"><span>&nbsp;&nbsp;&nbsp;Поел</span>
            </div>
        </div>
        <Auth ref="authPage"/>
        <SignUp ref="signUpPage"/>
    </div>
</template>
<script>
import Auth from "../auth/Auth.vue"
import SignUp from "../registration/Registration.vue"
export default {
    components: {
        Auth,
        SignUp,
    },

    methods: {
        showAuth: function() {
            this.$refs.authPage.show = true;
        },

        showSignUp: function() {
            this.$refs.signUpPage.show = true;
        }
    }
}
</script>