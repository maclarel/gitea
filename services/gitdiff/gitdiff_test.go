	"code.gitea.io/gitea/models/db"
	diff := `diff --git a/newfile2 b/newfile2
	diff2 := `diff --git "a/A \\ B" "b/A \\ B"
	diff2a := `diff --git "a/A \\ B" b/A/B
	diff3 := `diff --git a/README.md b/README.md

	assert.NoError(t, diff.LoadComments(db.DefaultContext, issue, user))
	expected := `<span class="line"><span class="cl">		<span class="n">run</span><span class="added-code"><span class="o">(</span><span class="n">db</span></span><span class="o">)</span>
</span></span>`