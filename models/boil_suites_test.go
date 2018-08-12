// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Announcements", testAnnouncements)
	t.Run("Archives", testArchives)
	t.Run("Articles", testArticles)
	t.Run("BlogTags", testBlogTags)
	t.Run("Blogcategories", testBlogcategories)
	t.Run("Bloggers", testBloggers)
	t.Run("BloggerMoods", testBloggerMoods)
	t.Run("Bloggerfriends", testBloggerfriends)
	t.Run("Comments", testComments)
	t.Run("Dtproperties", testDtproperties)
	t.Run("Events", testEvents)
	t.Run("Infos", testInfos)
	t.Run("Jokes", testJokes)
	t.Run("Links", testLinks)
	t.Run("Moods", testMoods)
	t.Run("Myhomes", testMyhomes)
	t.Run("Pics", testPics)
	t.Run("Pictures", testPictures)
	t.Run("Preferences", testPreferences)
	t.Run("Skins", testSkins)
	t.Run("TS", testTS)
	t.Run("Userdefinecategories", testUserdefinecategories)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Announcements", testAnnouncementsDelete)
	t.Run("Archives", testArchivesDelete)
	t.Run("Articles", testArticlesDelete)
	t.Run("BlogTags", testBlogTagsDelete)
	t.Run("Blogcategories", testBlogcategoriesDelete)
	t.Run("Bloggers", testBloggersDelete)
	t.Run("BloggerMoods", testBloggerMoodsDelete)
	t.Run("Bloggerfriends", testBloggerfriendsDelete)
	t.Run("Comments", testCommentsDelete)
	t.Run("Dtproperties", testDtpropertiesDelete)
	t.Run("Events", testEventsDelete)
	t.Run("Infos", testInfosDelete)
	t.Run("Jokes", testJokesDelete)
	t.Run("Links", testLinksDelete)
	t.Run("Moods", testMoodsDelete)
	t.Run("Myhomes", testMyhomesDelete)
	t.Run("Pics", testPicsDelete)
	t.Run("Pictures", testPicturesDelete)
	t.Run("Preferences", testPreferencesDelete)
	t.Run("Skins", testSkinsDelete)
	t.Run("TS", testTSDelete)
	t.Run("Userdefinecategories", testUserdefinecategoriesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Announcements", testAnnouncementsQueryDeleteAll)
	t.Run("Archives", testArchivesQueryDeleteAll)
	t.Run("Articles", testArticlesQueryDeleteAll)
	t.Run("BlogTags", testBlogTagsQueryDeleteAll)
	t.Run("Blogcategories", testBlogcategoriesQueryDeleteAll)
	t.Run("Bloggers", testBloggersQueryDeleteAll)
	t.Run("BloggerMoods", testBloggerMoodsQueryDeleteAll)
	t.Run("Bloggerfriends", testBloggerfriendsQueryDeleteAll)
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("Dtproperties", testDtpropertiesQueryDeleteAll)
	t.Run("Events", testEventsQueryDeleteAll)
	t.Run("Infos", testInfosQueryDeleteAll)
	t.Run("Jokes", testJokesQueryDeleteAll)
	t.Run("Links", testLinksQueryDeleteAll)
	t.Run("Moods", testMoodsQueryDeleteAll)
	t.Run("Myhomes", testMyhomesQueryDeleteAll)
	t.Run("Pics", testPicsQueryDeleteAll)
	t.Run("Pictures", testPicturesQueryDeleteAll)
	t.Run("Preferences", testPreferencesQueryDeleteAll)
	t.Run("Skins", testSkinsQueryDeleteAll)
	t.Run("TS", testTSQueryDeleteAll)
	t.Run("Userdefinecategories", testUserdefinecategoriesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Announcements", testAnnouncementsSliceDeleteAll)
	t.Run("Archives", testArchivesSliceDeleteAll)
	t.Run("Articles", testArticlesSliceDeleteAll)
	t.Run("BlogTags", testBlogTagsSliceDeleteAll)
	t.Run("Blogcategories", testBlogcategoriesSliceDeleteAll)
	t.Run("Bloggers", testBloggersSliceDeleteAll)
	t.Run("BloggerMoods", testBloggerMoodsSliceDeleteAll)
	t.Run("Bloggerfriends", testBloggerfriendsSliceDeleteAll)
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("Dtproperties", testDtpropertiesSliceDeleteAll)
	t.Run("Events", testEventsSliceDeleteAll)
	t.Run("Infos", testInfosSliceDeleteAll)
	t.Run("Jokes", testJokesSliceDeleteAll)
	t.Run("Links", testLinksSliceDeleteAll)
	t.Run("Moods", testMoodsSliceDeleteAll)
	t.Run("Myhomes", testMyhomesSliceDeleteAll)
	t.Run("Pics", testPicsSliceDeleteAll)
	t.Run("Pictures", testPicturesSliceDeleteAll)
	t.Run("Preferences", testPreferencesSliceDeleteAll)
	t.Run("Skins", testSkinsSliceDeleteAll)
	t.Run("TS", testTSSliceDeleteAll)
	t.Run("Userdefinecategories", testUserdefinecategoriesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Announcements", testAnnouncementsExists)
	t.Run("Archives", testArchivesExists)
	t.Run("Articles", testArticlesExists)
	t.Run("BlogTags", testBlogTagsExists)
	t.Run("Blogcategories", testBlogcategoriesExists)
	t.Run("Bloggers", testBloggersExists)
	t.Run("BloggerMoods", testBloggerMoodsExists)
	t.Run("Bloggerfriends", testBloggerfriendsExists)
	t.Run("Comments", testCommentsExists)
	t.Run("Dtproperties", testDtpropertiesExists)
	t.Run("Events", testEventsExists)
	t.Run("Infos", testInfosExists)
	t.Run("Jokes", testJokesExists)
	t.Run("Links", testLinksExists)
	t.Run("Moods", testMoodsExists)
	t.Run("Myhomes", testMyhomesExists)
	t.Run("Pics", testPicsExists)
	t.Run("Pictures", testPicturesExists)
	t.Run("Preferences", testPreferencesExists)
	t.Run("Skins", testSkinsExists)
	t.Run("TS", testTSExists)
	t.Run("Userdefinecategories", testUserdefinecategoriesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Announcements", testAnnouncementsFind)
	t.Run("Archives", testArchivesFind)
	t.Run("Articles", testArticlesFind)
	t.Run("BlogTags", testBlogTagsFind)
	t.Run("Blogcategories", testBlogcategoriesFind)
	t.Run("Bloggers", testBloggersFind)
	t.Run("BloggerMoods", testBloggerMoodsFind)
	t.Run("Bloggerfriends", testBloggerfriendsFind)
	t.Run("Comments", testCommentsFind)
	t.Run("Dtproperties", testDtpropertiesFind)
	t.Run("Events", testEventsFind)
	t.Run("Infos", testInfosFind)
	t.Run("Jokes", testJokesFind)
	t.Run("Links", testLinksFind)
	t.Run("Moods", testMoodsFind)
	t.Run("Myhomes", testMyhomesFind)
	t.Run("Pics", testPicsFind)
	t.Run("Pictures", testPicturesFind)
	t.Run("Preferences", testPreferencesFind)
	t.Run("Skins", testSkinsFind)
	t.Run("TS", testTSFind)
	t.Run("Userdefinecategories", testUserdefinecategoriesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Announcements", testAnnouncementsBind)
	t.Run("Archives", testArchivesBind)
	t.Run("Articles", testArticlesBind)
	t.Run("BlogTags", testBlogTagsBind)
	t.Run("Blogcategories", testBlogcategoriesBind)
	t.Run("Bloggers", testBloggersBind)
	t.Run("BloggerMoods", testBloggerMoodsBind)
	t.Run("Bloggerfriends", testBloggerfriendsBind)
	t.Run("Comments", testCommentsBind)
	t.Run("Dtproperties", testDtpropertiesBind)
	t.Run("Events", testEventsBind)
	t.Run("Infos", testInfosBind)
	t.Run("Jokes", testJokesBind)
	t.Run("Links", testLinksBind)
	t.Run("Moods", testMoodsBind)
	t.Run("Myhomes", testMyhomesBind)
	t.Run("Pics", testPicsBind)
	t.Run("Pictures", testPicturesBind)
	t.Run("Preferences", testPreferencesBind)
	t.Run("Skins", testSkinsBind)
	t.Run("TS", testTSBind)
	t.Run("Userdefinecategories", testUserdefinecategoriesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Announcements", testAnnouncementsOne)
	t.Run("Archives", testArchivesOne)
	t.Run("Articles", testArticlesOne)
	t.Run("BlogTags", testBlogTagsOne)
	t.Run("Blogcategories", testBlogcategoriesOne)
	t.Run("Bloggers", testBloggersOne)
	t.Run("BloggerMoods", testBloggerMoodsOne)
	t.Run("Bloggerfriends", testBloggerfriendsOne)
	t.Run("Comments", testCommentsOne)
	t.Run("Dtproperties", testDtpropertiesOne)
	t.Run("Events", testEventsOne)
	t.Run("Infos", testInfosOne)
	t.Run("Jokes", testJokesOne)
	t.Run("Links", testLinksOne)
	t.Run("Moods", testMoodsOne)
	t.Run("Myhomes", testMyhomesOne)
	t.Run("Pics", testPicsOne)
	t.Run("Pictures", testPicturesOne)
	t.Run("Preferences", testPreferencesOne)
	t.Run("Skins", testSkinsOne)
	t.Run("TS", testTSOne)
	t.Run("Userdefinecategories", testUserdefinecategoriesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Announcements", testAnnouncementsAll)
	t.Run("Archives", testArchivesAll)
	t.Run("Articles", testArticlesAll)
	t.Run("BlogTags", testBlogTagsAll)
	t.Run("Blogcategories", testBlogcategoriesAll)
	t.Run("Bloggers", testBloggersAll)
	t.Run("BloggerMoods", testBloggerMoodsAll)
	t.Run("Bloggerfriends", testBloggerfriendsAll)
	t.Run("Comments", testCommentsAll)
	t.Run("Dtproperties", testDtpropertiesAll)
	t.Run("Events", testEventsAll)
	t.Run("Infos", testInfosAll)
	t.Run("Jokes", testJokesAll)
	t.Run("Links", testLinksAll)
	t.Run("Moods", testMoodsAll)
	t.Run("Myhomes", testMyhomesAll)
	t.Run("Pics", testPicsAll)
	t.Run("Pictures", testPicturesAll)
	t.Run("Preferences", testPreferencesAll)
	t.Run("Skins", testSkinsAll)
	t.Run("TS", testTSAll)
	t.Run("Userdefinecategories", testUserdefinecategoriesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Announcements", testAnnouncementsCount)
	t.Run("Archives", testArchivesCount)
	t.Run("Articles", testArticlesCount)
	t.Run("BlogTags", testBlogTagsCount)
	t.Run("Blogcategories", testBlogcategoriesCount)
	t.Run("Bloggers", testBloggersCount)
	t.Run("BloggerMoods", testBloggerMoodsCount)
	t.Run("Bloggerfriends", testBloggerfriendsCount)
	t.Run("Comments", testCommentsCount)
	t.Run("Dtproperties", testDtpropertiesCount)
	t.Run("Events", testEventsCount)
	t.Run("Infos", testInfosCount)
	t.Run("Jokes", testJokesCount)
	t.Run("Links", testLinksCount)
	t.Run("Moods", testMoodsCount)
	t.Run("Myhomes", testMyhomesCount)
	t.Run("Pics", testPicsCount)
	t.Run("Pictures", testPicturesCount)
	t.Run("Preferences", testPreferencesCount)
	t.Run("Skins", testSkinsCount)
	t.Run("TS", testTSCount)
	t.Run("Userdefinecategories", testUserdefinecategoriesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Announcements", testAnnouncementsHooks)
	t.Run("Archives", testArchivesHooks)
	t.Run("Articles", testArticlesHooks)
	t.Run("BlogTags", testBlogTagsHooks)
	t.Run("Blogcategories", testBlogcategoriesHooks)
	t.Run("Bloggers", testBloggersHooks)
	t.Run("BloggerMoods", testBloggerMoodsHooks)
	t.Run("Bloggerfriends", testBloggerfriendsHooks)
	t.Run("Comments", testCommentsHooks)
	t.Run("Dtproperties", testDtpropertiesHooks)
	t.Run("Events", testEventsHooks)
	t.Run("Infos", testInfosHooks)
	t.Run("Jokes", testJokesHooks)
	t.Run("Links", testLinksHooks)
	t.Run("Moods", testMoodsHooks)
	t.Run("Myhomes", testMyhomesHooks)
	t.Run("Pics", testPicsHooks)
	t.Run("Pictures", testPicturesHooks)
	t.Run("Preferences", testPreferencesHooks)
	t.Run("Skins", testSkinsHooks)
	t.Run("TS", testTSHooks)
	t.Run("Userdefinecategories", testUserdefinecategoriesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Announcements", testAnnouncementsInsert)
	t.Run("Announcements", testAnnouncementsInsertWhitelist)
	t.Run("Archives", testArchivesInsert)
	t.Run("Archives", testArchivesInsertWhitelist)
	t.Run("Articles", testArticlesInsert)
	t.Run("Articles", testArticlesInsertWhitelist)
	t.Run("BlogTags", testBlogTagsInsert)
	t.Run("BlogTags", testBlogTagsInsertWhitelist)
	t.Run("Blogcategories", testBlogcategoriesInsert)
	t.Run("Blogcategories", testBlogcategoriesInsertWhitelist)
	t.Run("Bloggers", testBloggersInsert)
	t.Run("Bloggers", testBloggersInsertWhitelist)
	t.Run("BloggerMoods", testBloggerMoodsInsert)
	t.Run("BloggerMoods", testBloggerMoodsInsertWhitelist)
	t.Run("Bloggerfriends", testBloggerfriendsInsert)
	t.Run("Bloggerfriends", testBloggerfriendsInsertWhitelist)
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("Dtproperties", testDtpropertiesInsert)
	t.Run("Dtproperties", testDtpropertiesInsertWhitelist)
	t.Run("Events", testEventsInsert)
	t.Run("Events", testEventsInsertWhitelist)
	t.Run("Infos", testInfosInsert)
	t.Run("Infos", testInfosInsertWhitelist)
	t.Run("Jokes", testJokesInsert)
	t.Run("Jokes", testJokesInsertWhitelist)
	t.Run("Links", testLinksInsert)
	t.Run("Links", testLinksInsertWhitelist)
	t.Run("Moods", testMoodsInsert)
	t.Run("Moods", testMoodsInsertWhitelist)
	t.Run("Myhomes", testMyhomesInsert)
	t.Run("Myhomes", testMyhomesInsertWhitelist)
	t.Run("Pics", testPicsInsert)
	t.Run("Pics", testPicsInsertWhitelist)
	t.Run("Pictures", testPicturesInsert)
	t.Run("Pictures", testPicturesInsertWhitelist)
	t.Run("Preferences", testPreferencesInsert)
	t.Run("Preferences", testPreferencesInsertWhitelist)
	t.Run("Skins", testSkinsInsert)
	t.Run("Skins", testSkinsInsertWhitelist)
	t.Run("TS", testTSInsert)
	t.Run("TS", testTSInsertWhitelist)
	t.Run("Userdefinecategories", testUserdefinecategoriesInsert)
	t.Run("Userdefinecategories", testUserdefinecategoriesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Announcements", testAnnouncementsReload)
	t.Run("Archives", testArchivesReload)
	t.Run("Articles", testArticlesReload)
	t.Run("BlogTags", testBlogTagsReload)
	t.Run("Blogcategories", testBlogcategoriesReload)
	t.Run("Bloggers", testBloggersReload)
	t.Run("BloggerMoods", testBloggerMoodsReload)
	t.Run("Bloggerfriends", testBloggerfriendsReload)
	t.Run("Comments", testCommentsReload)
	t.Run("Dtproperties", testDtpropertiesReload)
	t.Run("Events", testEventsReload)
	t.Run("Infos", testInfosReload)
	t.Run("Jokes", testJokesReload)
	t.Run("Links", testLinksReload)
	t.Run("Moods", testMoodsReload)
	t.Run("Myhomes", testMyhomesReload)
	t.Run("Pics", testPicsReload)
	t.Run("Pictures", testPicturesReload)
	t.Run("Preferences", testPreferencesReload)
	t.Run("Skins", testSkinsReload)
	t.Run("TS", testTSReload)
	t.Run("Userdefinecategories", testUserdefinecategoriesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Announcements", testAnnouncementsReloadAll)
	t.Run("Archives", testArchivesReloadAll)
	t.Run("Articles", testArticlesReloadAll)
	t.Run("BlogTags", testBlogTagsReloadAll)
	t.Run("Blogcategories", testBlogcategoriesReloadAll)
	t.Run("Bloggers", testBloggersReloadAll)
	t.Run("BloggerMoods", testBloggerMoodsReloadAll)
	t.Run("Bloggerfriends", testBloggerfriendsReloadAll)
	t.Run("Comments", testCommentsReloadAll)
	t.Run("Dtproperties", testDtpropertiesReloadAll)
	t.Run("Events", testEventsReloadAll)
	t.Run("Infos", testInfosReloadAll)
	t.Run("Jokes", testJokesReloadAll)
	t.Run("Links", testLinksReloadAll)
	t.Run("Moods", testMoodsReloadAll)
	t.Run("Myhomes", testMyhomesReloadAll)
	t.Run("Pics", testPicsReloadAll)
	t.Run("Pictures", testPicturesReloadAll)
	t.Run("Preferences", testPreferencesReloadAll)
	t.Run("Skins", testSkinsReloadAll)
	t.Run("TS", testTSReloadAll)
	t.Run("Userdefinecategories", testUserdefinecategoriesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Announcements", testAnnouncementsSelect)
	t.Run("Archives", testArchivesSelect)
	t.Run("Articles", testArticlesSelect)
	t.Run("BlogTags", testBlogTagsSelect)
	t.Run("Blogcategories", testBlogcategoriesSelect)
	t.Run("Bloggers", testBloggersSelect)
	t.Run("BloggerMoods", testBloggerMoodsSelect)
	t.Run("Bloggerfriends", testBloggerfriendsSelect)
	t.Run("Comments", testCommentsSelect)
	t.Run("Dtproperties", testDtpropertiesSelect)
	t.Run("Events", testEventsSelect)
	t.Run("Infos", testInfosSelect)
	t.Run("Jokes", testJokesSelect)
	t.Run("Links", testLinksSelect)
	t.Run("Moods", testMoodsSelect)
	t.Run("Myhomes", testMyhomesSelect)
	t.Run("Pics", testPicsSelect)
	t.Run("Pictures", testPicturesSelect)
	t.Run("Preferences", testPreferencesSelect)
	t.Run("Skins", testSkinsSelect)
	t.Run("TS", testTSSelect)
	t.Run("Userdefinecategories", testUserdefinecategoriesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Announcements", testAnnouncementsUpdate)
	t.Run("Archives", testArchivesUpdate)
	t.Run("Articles", testArticlesUpdate)
	t.Run("BlogTags", testBlogTagsUpdate)
	t.Run("Blogcategories", testBlogcategoriesUpdate)
	t.Run("Bloggers", testBloggersUpdate)
	t.Run("BloggerMoods", testBloggerMoodsUpdate)
	t.Run("Bloggerfriends", testBloggerfriendsUpdate)
	t.Run("Comments", testCommentsUpdate)
	t.Run("Dtproperties", testDtpropertiesUpdate)
	t.Run("Events", testEventsUpdate)
	t.Run("Infos", testInfosUpdate)
	t.Run("Jokes", testJokesUpdate)
	t.Run("Links", testLinksUpdate)
	t.Run("Moods", testMoodsUpdate)
	t.Run("Myhomes", testMyhomesUpdate)
	t.Run("Pics", testPicsUpdate)
	t.Run("Pictures", testPicturesUpdate)
	t.Run("Preferences", testPreferencesUpdate)
	t.Run("Skins", testSkinsUpdate)
	t.Run("TS", testTSUpdate)
	t.Run("Userdefinecategories", testUserdefinecategoriesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Announcements", testAnnouncementsSliceUpdateAll)
	t.Run("Archives", testArchivesSliceUpdateAll)
	t.Run("Articles", testArticlesSliceUpdateAll)
	t.Run("BlogTags", testBlogTagsSliceUpdateAll)
	t.Run("Blogcategories", testBlogcategoriesSliceUpdateAll)
	t.Run("Bloggers", testBloggersSliceUpdateAll)
	t.Run("BloggerMoods", testBloggerMoodsSliceUpdateAll)
	t.Run("Bloggerfriends", testBloggerfriendsSliceUpdateAll)
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("Dtproperties", testDtpropertiesSliceUpdateAll)
	t.Run("Events", testEventsSliceUpdateAll)
	t.Run("Infos", testInfosSliceUpdateAll)
	t.Run("Jokes", testJokesSliceUpdateAll)
	t.Run("Links", testLinksSliceUpdateAll)
	t.Run("Moods", testMoodsSliceUpdateAll)
	t.Run("Myhomes", testMyhomesSliceUpdateAll)
	t.Run("Pics", testPicsSliceUpdateAll)
	t.Run("Pictures", testPicturesSliceUpdateAll)
	t.Run("Preferences", testPreferencesSliceUpdateAll)
	t.Run("Skins", testSkinsSliceUpdateAll)
	t.Run("TS", testTSSliceUpdateAll)
	t.Run("Userdefinecategories", testUserdefinecategoriesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Announcements", testAnnouncementsUpsert)
	t.Run("Archives", testArchivesUpsert)
	t.Run("Articles", testArticlesUpsert)
	t.Run("BlogTags", testBlogTagsUpsert)
	t.Run("Blogcategories", testBlogcategoriesUpsert)
	t.Run("Bloggers", testBloggersUpsert)
	t.Run("BloggerMoods", testBloggerMoodsUpsert)
	t.Run("Bloggerfriends", testBloggerfriendsUpsert)
	t.Run("Comments", testCommentsUpsert)
	t.Run("Dtproperties", testDtpropertiesUpsert)
	t.Run("Events", testEventsUpsert)
	t.Run("Infos", testInfosUpsert)
	t.Run("Jokes", testJokesUpsert)
	t.Run("Links", testLinksUpsert)
	t.Run("Moods", testMoodsUpsert)
	t.Run("Myhomes", testMyhomesUpsert)
	t.Run("Pics", testPicsUpsert)
	t.Run("Pictures", testPicturesUpsert)
	t.Run("Preferences", testPreferencesUpsert)
	t.Run("Skins", testSkinsUpsert)
	t.Run("TS", testTSUpsert)
	t.Run("Userdefinecategories", testUserdefinecategoriesUpsert)
	t.Run("Users", testUsersUpsert)
}
