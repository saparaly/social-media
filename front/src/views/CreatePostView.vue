<template>
    <div class="space"></div>

    <section class="createPost">
        <div class="container">
            <div class="createPost__container">
                <h2 class="title">Create Post</h2>
                <div class="error">{{ errorMess }}</div>
                <form action="">
                    <div class="form__group">
                        <input type="text" placeholder="title" name="title" v-model="title">
                    </div>
                    <div class="form__group">
                        <textarea id="" cols="30" rows="10" placeholder="description" name="description" v-model="description"></textarea>
                    </div>
                    <!-- add -->
                    <div class="add img" @click="add" v-if="!addfield">add img</div>
                    <div class="form__group"  v-if="addfield">
                        <input type="file" v-on:change="previewFile" name="img">
                        <div id="preview" v-html="previewContent"></div>
                        <div class="add" @click="add">cancel</div>
                    </div>
                    <div class="form__group">
                        <div v-for="(tag, index) in tags" :key="index" class="flex">
                            <input type="text" name="'tags" :value="tag" @input="updateTag($event, index, tag)">
                            <button @click.prevent="removeTag(index)" class="add">Remove</button>
                        </div>
                        <button class="add" @click.prevent="addTag()">Add Tag</button>
                    </div>

                    <div class="add" @click="add" v-if="!addfield">add click</div>
                    <div class="add" v-if="addfield">
                        <div class="form__group">
                            <input type="text" placeholder="date" name="date" v-model="date">
                        </div>
                        <div class="form__group">
                            <input type="text" placeholder="location" name="location" v-model="location">
                        </div>
                        <div class="add" @click="add">cancel</div>
                    </div>
                    <div class="form__group">
                        <input type="submit" value="create" class="create" @click.prevent="createPost">
                    </div>
                </form>
                 <div>
                 </div>
            </div>
        </div>
    </section>
</template>

<script>
import axios from 'axios'

export default {
    name: 'CreatePost',
    data() {
        return {
            previewContent: '',
            title: '',
            description: '',
            img: '',
            tags: [],
            created: new Date(),
            date: '',
            location: '',
            errorMess: '',
            addfield: false
        }
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
        async createPost() {
            try {
                const response = await axios.post('http://localhost:8000/create-post', {
                    Title: this.title,
                    Description: this.description,
                    Img: this.img,
                    Tags: this.tags,
                    Created: this.created,
                    Date: this.date,
                    Location: this.location
                },{
                    withCredentials: true,
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                console.log(response.data.valid)
                console.log(response.data)
                if (response.status === 200 && response.data.valid) {
                    this.$router.push(`/post/${response.data.postid}`)
                } else{
                    this.errorMess = response.data.err
                }
            } catch (error) {
                console.log(error)
            }
        },
        addTag() {
            this.tags.push('')
        },
        removeTag(index) {
            this.tags.splice(index, 1)
        },
        updateTag(event, index, tag) {
            this.tags[index] = event.target.value
        },
        add(){
            this.addfield = !this.addfield
        }
    }
}
</script>

<style scoped>
.form__group {
    margin-bottom: 20px;
}
.form__group input, .form__group textarea {
    border: 1px solid #8294C4;;
    border-radius: 5px;
    padding: 5px;
    resize: vertical;
    width: 100%;
    outline: none;
}

.add {
    padding: 5px 10px;
    cursor: pointer;
    background: #8294C4;;
    color: #fff;
    width: fit-content;
    border-radius: 5px;
    border: none;
}
.add.img {
    margin-bottom: 20px;
}
.create {
    padding: 10px;
    cursor: pointer;
    background: #8294C4;;
    color: #fff;
    border-radius: 5px;
    border: none;
    margin-top: 20px;
}
.flex {
    display: flex;
}
</style>