export default function ({ $axios }) {
  $axios.onRequest(config => {
    config.headers["token"] = window.localStorage.getItem("token")
  })

  $axios.onResponse(response => {
    localStorage.setItem("token", response.headers["token"])
  })
}
