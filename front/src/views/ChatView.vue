<template>
    <div style="height: 86vh">
        <div v-if="username && secret" style="height: 100%">
            <PrettyChatWindow
                :projectId="projectId"
                :username="username"
                :secret="secret"
            />
        </div>
    </div>
  </template>
  
  <style>
  .ce-new-chat-button {
    width: 32px;
    position: relative;
    bottom: 22px;
  }
  </style>
  
<script>
import axios from 'axios'

    import { PrettyChatWindow } from "react-chat-engine-pretty";
    import { applyReactInVue } from "veaury";
    export default {
        data() {
            return {
                projectId: import.meta.env.VITE_CHAT_ENGINE_PROJECT_ID,
                username: '',
                secret: '',
            }
        },
        components: {
            PrettyChatWindow: applyReactInVue(PrettyChatWindow),
        },
        async mounted() {
            try {
                const response = await axios.get('http://localhost:8000/profile', {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json'
                }
                });
                this.username = response.data.user.username
                this.secret = response.data.user.password
            } catch (error) {
                console.log(error);
            }
        },

    };
</script>