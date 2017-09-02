package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("아무런 값도 제공되지 않을 때, AuthAvatar.GetAvatarURL 은 ErrNoAvatarURL 를 리턴해야한다")
	}
	testUrl := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AtuhAvatar.GetAvatarURL 은 값이 제공되면 아무런 에러도 반환하지 않아야한다")
	}
	if url != testUrl {
		t.Error("AuthAvatar.GetAvatarURL 은 올바른 URL을 리턴해야한다")
	}
}
