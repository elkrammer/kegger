import axios from "axios";
import store from "@/store/";

axios.defaults.baseURL = config.API_SERVER;


axios.interceptors.response.use((response) => { return response; },
  async (error) => {
    const originalConfig = error.config;
    if (error.response) {
      if (error.response.status === 401 && !originalConfig._retry) {
        originalConfig._retry = true;
        store.dispatch("user/refreshToken")
        return axios(config);
      }

      // if (error.response.status === ANOTHER_STATUS_CODE) {
      //   // Do something 
      //   return Promise.reject(error.response.data);
      // }
    }

    return Promise.reject(error);
  }
);





// axios.interceptors.response.use(function (response) {
//     return response
// }, function (error) {
//     const { config, response: { status } } = error
//     const originalRequest = config
// 
//     if (status === 401 && error.response.data.message === "invalid or expired jwt") {
//         console.log("expired jwt, attempting to renew - axios.js")
//         store.dispatch("user/refreshToken")
//         let access_token = localStorage.getItem('accessToken');
// 
//         const retryOriginalRequest = new Promise((resolve) => {
//           originalRequest.headers.Authorization = 'Bearer ' + access_token
//           resolve(axios(originalRequest));
//         })
//         return retryOriginalRequest
//     }
//     return Promise.reject(error)
// })

export default axios;
