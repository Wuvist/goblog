-- 从原来sql server导入数据到mysql后，需要做的表结构修改，以符合sqlboiler需要

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
