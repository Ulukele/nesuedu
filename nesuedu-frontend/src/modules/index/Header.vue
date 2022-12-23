<template>
    <div class="header">
        <router-link to="/tasks" style="margin-left: 1.5%;">
            <div class="logo head">
                <span style="color:#434343; line-height: 200%;">NeSU.</span><span style="color:#87C644">Еду</span><span style="color:#9ED5F0">*</span>
            </div> 
        </router-link>
        <div class="acc-item" :style="{'background-color': nameColor}" @click="showUserPage"> {{ firstChar }} </div>
    </div>
</template>

<script>
import { axios_utils } from '../../axios_utils'

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
    data: () => ({
        nameColor: '',
        firstChar: '',
    }),

    mounted() {
        var userId = sessionStorage.getItem("id");
        var sessionID = sessionStorage.getItem("jwt");
        axios_utils.getUser(userId, sessionID).then(result => {
            this.firstChar = result.data.lastName.charAt(0);
            var surname = result.data.lastName;

            this.nameColor = '#' + intToRGB(hashCode(surname));
        })
    },

    methods: {
        showUserPage: function() {
            this.$emit('showUserPage', sessionStorage.getItem("id"));
        }
    }
}
</script>