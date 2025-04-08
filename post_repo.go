package main

func (p *postRepo) createPost(post *Post) (*Post, error) {
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
