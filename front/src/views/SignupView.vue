<template>
    <div class="signin">
      <div class="container">
        <div class="signin__container">
          <div class="left">
            <h1 class="title">Sign Up</h1>
            <p>Connect with Your Peers and Start Your Academic Adventure Now</p>
          </div>
          <div class="center">
            <img src="../assets/img/login.png" alt="">
          </div>
          <div class="rigth">
            <div class="error">{{ errorMess }}</div>
            <form action="">
              <div class="form__froup">
                <input type="file" v-on:change="previewFile" name="img">
                <div id="preview" v-html="previewContent"></div>
              </div>
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
                <div v-if="!passwordsMatch">Passwords do not match</div>
              </div>
              <div class="form__froup">
                <input type="password" name="password" placeholder="password" v-model="password">
              </div>
              <div class="form__froup">
                <input type="password" placeholder="confirm password" v-model="confirmPassword">
              </div>
              <div class="form__froup">
                <input type="submit" class="btn" value="submit" @click.prevent="signup">
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
import { loginRest, signupRest } from "./api";


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
        img: '',

        first_name: "",
        last_name: "",
        previewContent: '',

      }
    },
    mounted() {
  this.first_name = this.username;
  this.last_name = this.role;
},
  methods: {
    previewFile(event) {
            const file = event.target.files[0];
            const reader = new FileReader();
            reader.addEventListener('load', () => {
                const fileType = file.type.split('/')[0];
                if (fileType === 'image') {
                    this.previewContent = `<img src="${reader.result}" alt="preview">`;
                } else if (fileType === 'video') {
                    this.previewContent = `<video src="${reader.result}" controls></video>`;
                } else {
                    this.previewContent = '<p>Invalid file type.</p>';
                }
                this.img = reader.result; 
            });
            reader.readAsDataURL(file);
        },
    async signup() {
      try {
        const response = await axios.post('http://localhost:8000/signup', {
          avaimg: this.img,
          role: this.role,
          username: this.username,
          email: this.email,
          password: this.password,
        })
        console.log(response.data)
        if (response.status === 200 && response.data.valid) {
          // redirect to signin page
          this.chatSignup();
          this.$router.push('/signin')
        } else{
          this.errorMess = response.data.err
        }

      } catch (error) {
        console.log(error)
      }
    },
    chatSignup() {
        signupRest(
          this.username,
          this.password,
          this.email,
          this.username,
          this.role
        )
          .then((response) =>
            this.$emit("onAuth", { ...response.data, secret: this.password })
          )
          .catch((error) => console.log("Sign up error", error));
      },
  },
  computed: {
    passwordsMatch() {
      return this.password === this.confirmPassword
    }
  }
}
</script>

<style scoped>
.img {
  display: none;
}

.form__froup {
  margin-bottom: 20px;
}
.form__froup input {
  width: 100%;
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
  width: 400px;
}
.title{
  font-weight: 600;
  font-size: 58px;
  color: #404040;
  margin-bottom: 30px;
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
</style>
  