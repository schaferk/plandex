package shared

func (f *ConvoMessageFlags) GetReplyTags() []string {
	var replyTags []string
	if f.DidLoadContext {
		replyTags = append(replyTags, "📥 Loaded Context")
	}
	if f.DidMakePlan {
		if f.DidMakeDebuggingPlan {
			replyTags = append(replyTags, "🐞 Made Debug Plan")
		} else {
			replyTags = append(replyTags, "📋 Made Plan")
		}
	}
	if f.DidWriteCode {
		replyTags = append(replyTags, "👨‍💻 Wrote Code")
	}
	// if f.DidCompleteTask {
	// 	replyTags = append(replyTags, "✅")
	// }
	if f.DidCompletePlan {
		replyTags = append(replyTags, "🏁")
	}

	return replyTags
}
