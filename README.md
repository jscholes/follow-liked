# follow-liked

The Release Radar playlist on Spotify is regenerated weekly, and seems to primarily contain tracks followed by the current user.  There are some forum posts and other content around the web indicating that it takes other aspects of the user's listening profile into account, but I consistently find that tracks bubble up (and/or take the place of other ones I would enjoy more) from artists that I followed years ago.

As such, this is an experiment to see if following the artists I actually enjoy today will improve the utility of Release Radar.  I'm not in the habit of following artists when I like one of their songs, and wish Spotify would do so automatically.  But it doesn't, so this is a small command line utility that will:

1. fetch all of your liked songs;
2. work through the artists for each of those songs, ignoring any duplicates; and
3. follow each artist.

It will also store a list of followed artists in the same directory that you run the tool from, which will include:

* any artist previously explicitly followed through the use of this tool; and
* artists that you follow already, fetched from your user profile.

As a future enhancement, I'd like it to provide a `-u`/`--unfollow` switch that allows a specific artist to be unfollowed by ID (e.g. to clean out artists that you no longer care about).