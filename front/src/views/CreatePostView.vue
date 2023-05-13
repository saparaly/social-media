<template>
    <!-- <div class="space"></div> -->

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
                    <div class="add img" @click="addimg" v-if="!addfield">add img</div>
                    <div class="form__group"  v-if="addfield">
                        <input type="file" v-on:change="previewFile" name="img">
                        <div id="preview" v-html="previewContent"></div>
                        <div class="add mt" @click="addimg">cancel</div>
                    </div>
                    <div class="form__group">
                        <div v-for="(tag, index) in tags" :key="index" class="flex mb">
                            <input type="text" name="'tags" :value="tag" @input="updateTag($event, index, tag)">
                            <button @click.prevent="removeTag(index)" class="add">Remove</button>
                        </div>
                        <button class="add" @click.prevent="addTag()">Add Tag</button>
                    </div>

                    <div class="add" @click="adddate" v-if="!addfieldd">add date</div>
                    <div v-if="addfieldd">
                        <div class="form__group">
                            <input type="text" placeholder="date" name="date" v-model="date">
                        </div>
                        <div class="form__group">
                            <input type="text" placeholder="location" name="location" v-model="location">
                        </div>
                        <div class="add" @click="adddate">cancel</div>
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
            addfieldd: false,
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
        addimg(){
            this.addfield = !this.addfield
        },
        adddate(){
            this.addfieldd = !this.addfieldd
        }
    }
}
</script>

<style scoped>
.img {
    width: 700px;
}
.img img {
    width: 100%;
    height: 100%;
    object-fit: contain;
}
.form__group, .mb {
    margin-bottom: 20px;
}
.mt{
    margin-top: 20px;
}
.form__group input, .form__group textarea {
    /* border: 1px solid var(--primary-color);;
    border-radius: 5px;
    padding: 5px;
    resize: vertical;
    width: 100%;
    outline: none; */
    border: 1px solid #676767;
    border-radius: 10px;
    padding: 10px 20px;
    font-size: 18px;
    letter-spacing: 0.09em;
    color: #4F555A;
    background: transparent;
    outline: #4F555A;
    width: 100%;
}

.add {
    padding: 5px 10px;
    cursor: pointer;
    background: var(--primary-color);
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
    background: var(--primary-color);
    color: #fff;
    border-radius: 5px;
    border: none;
    margin-top: 20px;
}
.flex {
    display: flex;
}
</style>