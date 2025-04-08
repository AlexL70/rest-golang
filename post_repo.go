package main

func (p *postRepo) createPost(post *Post) (*Post, error) {
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (p *postRepo) getPost(id uint64) (*Post, error) {
	var post = new(Post)
	if err := p.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}
