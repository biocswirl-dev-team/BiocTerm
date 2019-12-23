# BiocTerm

A highly configurable, lightweight and feature rich terminal interface with multi-language support for conducting bioinformatics. This project was the People's Choice Award for the Vancouver Bioinformatics Hackathon Hackseq2019.

BiocTerm is a standalone terminal application that acts as the ideal interface for those conducting bioinformatics as it integrates our custom editor, biovim (a vim installation bundled with bio plugins), radian (an r console), and tmux (our customized terminal wrapper) in a contained instance. Our interface supports R, Bash, and vim. It is ideal for multi language workflows.

<img scr="https://cdn.discordapp.com/attachments/658422882312912928/658472806153060352/unknown.png">

### Background

(In progress) 

### Features

- Automatically saves session info + restores and generates a log 
- Multiple panes for native file viewing, has robust syntax highlighting that is ideal for bioinformaticians such as BioSyntax or Radian
- Tmux.conf file and plugin capabilities for vim and tmux allow for high configurability depending on what your needs are 
- Comes with our editor biovim, a powerful editor that makes reading gene files easy

### Examples of use: (screenshots) 


Please also see our bioinformatics course material that runs natively through R, ![biocswirl](https://github.com/biocswirl-dev-team/BiocSwirl). The interface (BiocTerm) and R package (BiocSwirl) can be used independently of each other but are best used together for people who are interested in making the most of their learning experience.


### Team members:

| Name | GitHub ID | Work done |
| ---- | --------- | --------- |
| Lisa N. Cao | lisancao | User Environment & Interface Development |
| Jackie Lu | jql6 | User Environment & Interface Development (tmux) |
| Jeremy Fan | zhemingfan | User Environment & Interface Development |
| Mariam Arab | mariamarab | User Environment & Installation |
| Kate Tyshchenko | ktyshchenko | Documentation, Course Material |
| Paaksum Wong | paaksum | Course Material (main) |
| Sourav Singh | souravsingh | Course Material |   

[Details about hackseq19 project](biocswirl_dev/hackseq_plan)


### Contributing

We are always looking for pull requests and active contributers, if you are interested in designing a course for us or have a feature in mind please submit an issue before doing a pull request. We are currently looking for help developing on Windows systems, bioinformatics workflows and concepts, and support for even more languages in BiocTerm. 
