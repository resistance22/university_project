import http from 'k6/http'

export let options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  stages: [
    {
      duration: '10s',
      target: 100
    },
    {
      duration: '10s',
      target: 100
    },
    {
      duration: '10s',
      target: 200
    },
    {
      duration: '10s',
      target: 300
    },
    {
      duration: '10s',
      target: 400
    },
  ]
}

export default () => {
  const payload = JSON.stringify({
    user_name: "amin_f",
    password: "complexP@ss"
  })

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post('http://localhost:8080/api/v1/auth/login', payload, params)
}