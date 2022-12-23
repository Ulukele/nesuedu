<template>
    <Transition name="slide-fade">
        <div v-if="show" class="reg-window" @click.self="closePage">
            <div class="auth-form">
                <div class="logo reg">
                    <span style="color:#434343">NeSU.</span><span style="color:#87C644">Еду</span><span style="color:#9ED5F0">*</span>
                </div>  

                <div class="form">          
                    <div class="reg-fields">
                        <va-input
                            v-model="firstName"
                            placeholder="Имя"
                            class="reg-input"
                            />

                        <va-input
                            v-model="lastName"
                            placeholder="Фамилия"
                            class="reg-input"
                            />

                        <va-input
                            v-model="email"
                            type="email"
                            placeholder="Почта"
                            class="reg-input"
                            :error="emailErrorMsg !== ''"
                            :error-messages="emailErrorMsg"
                            />

                        <va-input
                        v-model="password"
                        :type="isPasswordVisible ? 'text' : 'password'"
                        placeholder="Пароль"
                        class="reg-input"
                        :error="passwordErrorMsg !== ''"
                        :error-messages="passwordErrorMsg"
                        >
                        <template #appendInner>
                        <va-icon
                            :name="isPasswordVisible ? 'visibility_off' : 'visibility'"
                            size="medium"
                            color="#808080"
                            @click="isPasswordVisible = !isPasswordVisible"
                        />
                        </template>
                        </va-input>

                        <va-input
                        v-model="repeatPassword"
                        :type="isPasswordVisible ? 'text' : 'password'"
                        placeholder="Повторите пароль"
                        class="reg-input"
                        :error="passwordErrorMsg !== ''"
                        :error-messages="passwordErrorMsg"
                        >
                        <template #appendInner>
                        <va-icon
                            :name="isPasswordVisible ? 'visibility_off' : 'visibility'"
                            size="medium"
                            color="#808080"
                            @click="isPasswordVisible = !isPasswordVisible"
                        />
                        </template>
                        </va-input>

                        <div style="position: relative; display: flex; justify-content:space-between; width: 58%;">
                            <va-button 
                            :disabled="!(password !== '' && email !== '' && repeatPassword !== '' && firstName !== '' && lastName !== '' && password === repeatPassword)"
                            color="customGreen"
                            class="reg-btn"
                            @click="tryToSignUp"
                            >Sign Up</va-button>

                            <va-button 
                            color="customRed"
                            class="reg-btn"
                            @click="closePage"
                            >Cancel</va-button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Transition>
</template>

<script>
import { axios_utils } from '../../axios_utils'
export default {
    data: () => ({
        show: false,
        isPasswordVisible: false,
        passwordErrorMsg: '',
        emailErrorMsg: '',
        firstName: '',
        lastName: '',
        password: '',
        repeatPassword: '',
        email: '',
    }),

    methods: {
        tryToSignUp: function() {
            var data_ = {
                'firstName': this.firstName,
                'lastName': this.lastName,
                'username': this.email,
                'password': this.password,
            }
            axios_utils.tryToSignUp(data_).then(result => {
                sessionStorage.setItem("jwt", result.data.jwt);
                sessionStorage.setItem("id", result.data.id);
                sessionStorage.setItem("refreshToken", result.data.refreshToken);
                const routeData = this.$router.resolve({name: 'tasks'});
                window.open(routeData.href, '_self');
            });
        },

        closePage: function() {
            this.show = false;
            this.isPasswordVisible = false;
            this.password = '';
            this.repeatPassword = '';
            this.email = '';
            this.firstName = '';
            this.lastName = '';
            this.passwordErrorMsg = '';
            this.emailErrorMsg = '';
        }
    }
}
</script>