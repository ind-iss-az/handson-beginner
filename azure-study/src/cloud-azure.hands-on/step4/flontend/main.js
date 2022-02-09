const API_GATEWAY_URL = "YOUR_API_GATEWAY_URL"

var app = new Vue({
    el: '#app',

    data() {
        return {
            date: "",
            hostname: "",
            obj: null,
          }
    },

    mounted() {
        var config = {
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Credentials': 'true'
            }
        };

        axios.get(API_GATEWAY_URL + "/hello", config)
        .then(response => {
            obj = JSON.parse(response.data.body);
            this.obj = obj;
        })
        .catch(error => {
            console.log(error);
        });
    }
})

