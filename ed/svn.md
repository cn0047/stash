svn
-

````
svn st      # status
svn di file # diff
````

#### The typical work cycle looks like this:

````
# Check out a working copy from a repository.
svn checkout your_name https://svn.server.com/repository/trunk

# Update your working copy
svn update

# Make changes
svn add
svn delete
svn copy
svn move

# Examine your changes
svn status
svn diff
svn revert

# Merge others' changes
svn merge
svn resolved

# Commit your changes
svn commit
````
