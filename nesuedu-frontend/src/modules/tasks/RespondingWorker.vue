<template>
    <div style="height: fit-content; width:100%; display: flex; font-size: 1.2rem;">
        <div class="worker acc-item" :style="{'background-color': this.nameColor}" @click="showUserPage">
            {{ worker.lastName.charAt(0) }}
        </div>
        <div style="align-self: center; margin-left: 2.5%; min-width: 30%; font-weight: bolder; color: #434343;">
            {{ worker.firstName + ' ' + worker.lastName }}
        </div>

        <div class="rating" :style="{'color': this.rateColor}" >
            {{ worker.rate }}
        </div>

        <button class="accept-btn">Подтвердить</button>
    </div>
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
    props: ['worker'],

    data: () => ({
        nameColor: '',
        rateColor: '',
    }),

    beforeMount() {
        this.nameColor = '#' + intToRGB(hashCode(this.worker.lastName));

        if (this.worker.rate >= 4) {
            this.rateColor = '#87C644'; 
        } else if (this.worker.rate >= 3 && this.worker.rate < 4) {
            this.rateColor = '#D6DA04'; 
        } else if (this.worker.rate >= 2 && this.worker.rate < 3) {
            this.rateColor = '#FFB443'; 
        } else if (this.worker.rate >= 1 && this.worker.rate < 2) {
            this.rateColor = '#FF792E'; 
        } else if (this.worker.rate < 1) {
            this.rateColor = '#FF6A6A'; 
        } 
    },

    methods: {
        showUserPage: function() {
            this.$emit('showUserPage', this.worker.id);
        }
    }
}
</script>
