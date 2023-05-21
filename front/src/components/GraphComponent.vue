<template>
  <div>
    <h3 class="title">Bachelor Programmes Astana IT University</h3>
        <ul class="tab-nav">
          <li v-for="(tab, index) in tabs" :key="index" :class="{ active: activeTab === index }" @click="activeTab = index">
            {{ tab.title }}
          </li>
        </ul>
        <div class="tab-content">
          <div v-for="(tab, index) in tabs" :key="index" v-show="activeTab === index" class="tab">
            <div v-for="(con, i) in tab.content" :key="i" class="con">
                <div v-for="(val, ind) in con" :key="ind" class="val">
                    <div class="items">{{ val[0] }} <span>{{ val[1] }}</span></div>
                </div>
            </div>
          </div>
        </div>
  </div>
</template>

<script>
export default {
data() {
    return {
        activeTab: 0,
          tabs: [
            { 
                title: 'Software Engineering', 
                content:[
                     [
                        ['Term 1', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 2', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 3', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 4', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 5', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 6', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 7', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 8', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 9', 26],
                        ['Diploma Work (Project) and Defence', 12] ,
                        ['Undergraduate practice', 4],
                        ['Industrial practice', 8]
                    ],
                ]
            },
            { title: 'Computer Science', 
            content:[
                     [
                        ['Term 1', 29],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 2', 24],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 3', 27],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 4', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 5', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 6', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 7', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 8', 26],
                        ['Foreign Language (English 1)', 5] ,
                        ['Cultural Studies', 5],
                        ['Information and Communication Technologies', 5],
                        ['Calculus 1', 5],
                        ['Introduction to Programming', 5],
                        ['Physical Culture', 2], 
                        ['Educational Practice', 2]
                    ],
                     [
                        ['Term 9', 26],
                        ['Diploma Work (Project) and Defence', 12] ,
                        ['Undergraduate practice', 4],
                        ['Industrial practice', 8]
                    ],
                ]
            },
            { title: 'Big Data', content: '' },
            { title: 'Industrial Automation', content: '' },
            { title: 'Media Technologies', content: '' },
            { title: 'IT Management', content: '' },
            { title: 'Cybersecurity', content: '' },
            { title: 'Telecommunication', content: '' },
          ],
    }
}
}
</script>

<style scoped>
.items{
    color: var(--dark);
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 20px;
    border: 1px solid;
    padding: 10px;
    border-radius: 10px;
    letter-spacing: 0.09em;
    color: #4F555A;
    margin-bottom: 20px;
}
.items span{
    color: var(--text-color);
    font-weight: bold;
}
.val:first-child .items{
    background: var(--dark);
    color: #fff;
    font-weight: bold;
}
.val:first-child .items span{
    color: #fff;
}
.tab {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 30px;
}
.con {
    box-shadow: 0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04);
    background: #fff;
    border-radius: 20px;
    padding: 20px;
}
.tab-nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 30px;
  margin-bottom: 30px;
}

.tab-nav li {
    background: var(--dark);
    border: 2px solid var(--dark);
    color: #fff;
    padding: 5px;
    border-radius: 5px;
    cursor: pointer;
}

.tab-nav li.active {
    background: #fff;
    color: var(--dark);
    font-weight: bold;
}
</style>