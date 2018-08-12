ALTER TABLE `blogwind`.`users` 
ADD PRIMARY KEY (`id`);
;


ALTER TABLE `blogwind`.`customer` 
ADD PRIMARY KEY (`id`);
;

ALTER TABLE `blogwind`.`bloggerfriend` 
ADD PRIMARY KEY (`blogger`,`friend`);
;

ALTER TABLE `blogwind`.`pictures` 
ADD PRIMARY KEY (`id`);
;

ALTER TABLE `blogwind`.`skin` 
ADD PRIMARY KEY (`index`);
;
