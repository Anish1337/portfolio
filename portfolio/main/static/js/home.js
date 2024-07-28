// todo
document.addEventListener("DOMContentLoaded", function () {
    const nameSearch = document.getElementById("name-search")
    const tags = document.querySelectorAll(".tag")
    const projects = document.querySelectorAll(".project")

    function filterProjects() {
        const nameQuery = nameSearch.value.toLowerCase();
        
        // loop through each project
        projects.forEach((project) => {
            // get name of project
            const name = project.getAttribute('data-name')
            // boolean includes
            const nameMatch = name.includes(nameQuery)
            // if matches
            if (nameMatch) {
                project.style.display = "";
            } else
            // otherwise
            {
                project.style.display = "none";
            }
        })
    }

    tags.forEach((tag) => {
        tag.addEventListener("click", function () {
            const selectedTag = this.getAttribute("data-tag")

            projects.forEach((project) => {
                const projectTags = project.getAttribute("data-tags")
                if (projectTags.includes(selectedTag)) {
                    project.style.display = ""
                } else {
                    project.style.display = "none"
                }
            })
        })
    })
    
    // Whenever key is lifted call the search function
    nameSearch.addEventListener("keyup", filterProjects)
})
