package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL은 Avatar 인스턴스가 URL을 제공할 수 없을 때 제공되는 에러
var ErrNoAvatarURL = errors.New("chat: avatar URL을 가져올 수 없습니다")

// Avatar는 사용자 프로필 사진을 조회할 수 있는 타입
type Avatar interface {
	// GetAvatarURL은 지정된 클라이언트에 대한 아바타 URL을 가져오고, 문제가 발생하면 에러를 리턴
	// 객체가 지저된 클라이언트의 URL을 가져올 수 없는 경우, ErrNoAvatarURL을 리턴
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
