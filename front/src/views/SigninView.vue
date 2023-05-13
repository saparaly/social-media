<template>
  <div class="signin">
      <div class="container">
        <div class="signin__container">
          <div class="left">
            <div class="signup">
              <h2 class="title">Unlock Your Access, Sign Up Now</h2>
              <p>Join the Community, Join the Conversation</p>
              <RouterLink to="/signup">signup</RouterLink>
            </div>
          </div>
          <div class="center">
            <img src="../assets/img/login.png" alt="">
          </div>
          <div class="right">
          <div class="form__froup">
            <p>Stay Informed and Sign In to Your University Account</p>
            <div class="error">{{ errorMess }}</div>
          </div>
          <form action="" >
            <div class="form__froup">
              <input type="text" name="username" placeholder="username" v-model="username">
            </div>
            <div class="form__froup">
              <input type="password" name="password" placeholder="password" v-model="password">
            </div>
            <div class="form__froup">
              <input type="submit" value="submit" @click.prevent="signin">
            </div>
          </form>
          </div>
        </div>
        <div class="img"><img src="../assets/img/Background.png" alt=""></div>

      </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      username: '',
      password: '',
      errorMess: '',
    }
  },
methods: {
  async signin() {
    try {
      const response = await axios.post('http://localhost:8000/signin', {
        username: this.username,
        password: this.password,
      },{
        withCredentials: true
      })
      // con sole.log(response.data)
      console.log(response.data.valid)
        if (response.status === 200 && response.data.valid) {
          this.$router.push('/')
        } else{
          this.errorMess = response.data.err
        }
    } catch (error) {
      console.log(error)
    }
  }
},
}
</script>

<style scoped>
.img {
  display: none;
}
.error {
  color: red;
}

.form__froup {
  margin-bottom: 20px;
}
.form__froup input {
  border: 1px solid #676767;
  border-radius: 10px;
  padding: 10px 20px;
  font-size: 18px;
  letter-spacing: 0.09em;
  color: #4F555A;
  background: transparent;
  outline: #4F555A;
}
input.btn {
  width: 100%;
  background: #4461F2;
  box-shadow: 0px 12px 21px 4px rgba(68, 97, 242, 0.15);
  border-radius: 10px;
  border: none;
  font-weight: 600;
  letter-spacing: 0.09em;
  color: #FFFFFF;
  cursor: pointer;
  transition: 0.3s;
}
input.btn:hover {
  transform: scale(0.9);
}
.left {
  width: 600px;
}

p{
  font-size: 18px;
  letter-spacing: 0.09em;
  color: #4F555A;
}
.img{
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  z-index: -1;

}
.img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.signin {
  /* padding-bottom: 100px; */
}
.signin__container {
  display: flex;
  align-items: center;
}
/* .form__froup p {
  margin-bottom: 10px;
}
.test{
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}
.oneContainer{
position: fixed;
top: 50%;
left: 50%;
transform: translate(-50%, -50%);
width: 1000px;
height: 800px;
}
p {
color: #AB8E91;
font-weight: bold;
}
.signin {
  position: absolute;
  top: 50%;
  left: 0;
  transform: translate(0%, -50%);
  backdrop-filter: blur(10px);
  background: #fff;
  width: 450px;
  padding: 20px;
  border-radius: 20px;
  box-shadow:
      inset 0 -1em 1em rgba(0,0,0,0.1),
            0.1em 0.1em 1em rgba(0,0,0,0.3);
  text-align: center;   
  height: 500px;
  display: flex;        
  flex-direction: column;
  justify-content: center;
  align-items: center;
  z-index: 100;
}
.form__froup {
display: flex;
flex-direction: column;
margin-bottom: 10px;
}
.form__froup input {
border: none;
background: #485680;
border-radius: 15px;
padding: 15px;
color: #FFF;
}

.signup{
  position: absolute;
  top: 0;
  right: 0;
  backdrop-filter: blur(10px);
  width: 600px;
  padding: 20px;
  border-radius: 20px;
  box-shadow:
      inset 0 -1em 1em rgba(0,0,0,0.1),
            0.1em 0.1em 1em rgba(0,0,0,0.3);
  text-align: center;   
  height: 800px;
  display: flex;        
  flex-direction: column;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}
.signup img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  position: absolute;
  z-index: -1;

} */
</style>
