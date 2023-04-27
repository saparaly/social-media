<template>
    <div class="signin">
      <h1 class="title">Sign Up</h1>
      <p>Connect with Your Peers and Start Your Academic Adventure Now</p>
      <div class="error">{{ errorMess }}</div>
      <form action="" >
        <div class="form__froup">
          <input type="text" placeholder="role" name="role" v-model="role">
        </div>
        <div class="form__froup">
          <input type="text" name="username" placeholder="username" v-model="username">
        </div>
        <div class="form__froup">
          <input type="email" name="email" placeholder="email" v-model="email">
        </div>
        <div class="form__froup">
          <input type="password" name="password" placeholder="password" v-model="password">
        </div>
        <div class="form__froup">
          <input type="password" placeholder="confirm password" v-model="confirmPassword">
        </div>
        <div class="form__froup">
          <div v-if="!passwordsMatch">Passwords do not match</div>
        </div>
        <div class="form__froup">
          <input type="submit" value="submit" @click.prevent="signup">
        </div>
      </form>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
      return {
        user: {},
        role: '',
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
        errorMess: '',
      }
    },
  methods: {
    async signup() {
      try {
        const response = await axios.post('http://localhost:8000/signup', {
          role: this.role,
          username: this.username,
          email: this.email,
          password: this.password,
        })
        console.log(response.data)
        if (response.status === 200 && response.data.valid) {
          // redirect to signin page
          this.$router.push('/signin')
        } else{
          this.errorMess = response.data.err
        }

      } catch (error) {
        console.log(error)
      }
    }
  },
  computed: {
    passwordsMatch() {
      return this.password === this.confirmPassword
    }
  }
}
</script>

<style scoped>
p {
  color: #AB8E91;
  font-weight: bold;
}
.signin {
  backdrop-filter: blur(10px);
  width: 500px;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 30px;
  border-radius: 20px;
  box-shadow:
       inset 0 -1em 1em rgba(0,0,0,0.1),
             0.1em 0.1em 1em rgba(0,0,0,0.3);
  text-align: center;   
  height: fit-content;        
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

.cloud__img {
  position: fixed;
  z-index: -1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>
  