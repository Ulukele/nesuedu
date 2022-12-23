<script>
import { axios_utils } from '../../axios_utils'
import { VueElement } from 'vue'
export default {
  methods: {
    tryToRefreshToken() {
      console.log('try to refresh')
      var refreshToken = sessionStorage.getItem("refreshToken");
      if (!refreshToken) {
        return
      }
      console.log('refresh')
        axios_utils.refreshToken(refreshToken).then(
          result => {
            sessionStorage.setItem("refreshToken", result.data.refreshToken);
            sessionStorage.setItem("jwt", result.data.jwt);
            sessionStorage.setItem("id", result.data.id);
          },
          error => { console.error(error) }
        )
    },

  },
  mounted() {
		this.intervalId = setInterval(this.tryToRefreshToken, 3000 * 60);
    var refreshToken = sessionStorage.getItem("refreshToken");
    if (refreshToken) {
      this.tryToRefreshToken();
    }
  },

  beforeDestroy() {
		clearInterval(this.intervalId);
	},
}
</script>

<template>
  <router-view></router-view>
</template>