<template>
    <Transition name="slide-fade">
        <div v-if="show" class="reg-window" @click.self="closePage">
            <div class="auth-form" style="height: 57%">
                <div v-if="!isLoaded" class="loading-screen">
                    <va-progress-circle indeterminate 
                    color="#9ED5F0"
                    :thickness="0.3"
                    size="large"/>
                </div>
                <div v-else class="user-page">
                    <div style="position: relative; display: flex; gap: 20px; align-items: center; margin-bottom: 3.5%;">
                        <div class="acc-item acc-page" :style="{'background-color': this.nameColor}"> 
                            {{  user.lastName.charAt(0) }} 
                        </div>
                        <a href=""><img src="/images/vk-icon.svg"></a>
                        <a href=""><img src="/images/telegram-icon.svg"></a>
                    </div>
                    <div class="user-name">{{ user.firstName + ' ' + user.lastName }}</div>
                    <div class="user-email">{{ user.username }}</div>
                    <div class="rate-area space-betw"> 
                        <div style="justify-content: flex-start;">
                            Рейтинг: 
                        </div>
                        <div :style="{'color': this.rateColor}">
                            {{ user.rate }}
                        </div>
                        <div class="rate-finger" :style="{'--rate-deg': this.deg}">
                            👍
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Transition>
</template>

<script>

function hashCode(str) {
    var hash = 0;
    for (var i = 0; i < str.length; i++) {
       hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return hash;
} 

function intToRGB(i){
    var c = (i & 0x00FFFFFF)
        .toString(16)
        .toUpperCase();

    return "00000".substring(0, 6 - c.length) + c;
}

export default {
    props: ['user'],

    data: () => ({
        show: false,
        isLoaded: false,
        nameColor: '',
        rateColor: '',
        deg: '0deg',
    }),

    beforeUpdate() {
        this.setNameColor();
        this.setRateColor();
        this.setRateDeg();
    },

    methods: {
        closePage: function() {
            this.show = false;
        },

        setRateColor: function() {
            this.rateColor = '';
            if (this.user.rate >= 4) {
                this.rateColor = '#87C644'; 
            } else if (this.user.rate >= 3 && this.user.rate < 4) {
                this.rateColor = '#D6DA04'; 
            } else if (this.user.rate >= 2 && this.user.rate < 3) {
                this.rateColor = '#FFB443'; 
            } else if (this.user.rate >= 1 && this.user.rate < 2) {
                this.rateColor = '#FF792E'; 
            } else if (this.user.rate < 1) {
                this.rateColor = '#FF6A6A'; 
            } 
        },

        setNameColor: function() {
            this.nameColor = '#' + intToRGB(hashCode(this.user.lastName));
        },

        setRateDeg: function() {
            this.deg = '0deg';
            this.deg = (180 - 45 * (this.user.rate - 1)).toString() + 'deg';
        },
    }
}
</script>